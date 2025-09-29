package chapter12

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	c      = context.Background()
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
)

type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100"`
	Price int
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Product{})
	db.Create(&Product{Name: "Product1", Price: 1382})

	return db
}

func getProductByIdHash(db *gorm.DB, id uint) (Product, error) {
	cacheKey := fmt.Sprintf("product:%d", id)

	var product Product

	res, err := client.HGetAll(c, cacheKey).Result()
	if err == nil && len(res) > 0 {
		product.ID = id
		product.Name = res["name"]
		product.Price, _ = strconv.Atoi(res["price"])
		return product, nil
	}

	if err := db.First(&product, id).Error; err != nil {
		return product, err
	}

	client.HMSet(c, cacheKey, map[string]any{
		"name":  product.Name,
		"price": product.Price,
	})

	client.Expire(c, cacheKey, time.Minute)

	return product, nil
}

func addToRecentProductList(id uint) {
	client.LPush(c, "recent_products", id)
	client.LTrim(c, "recent_products", 0, 9)
}

func createOrUpdateProductWriteThrough(db *gorm.DB, id uint, name string, price int) error {
	product := Product{
		ID:    id,
		Name:  name,
		Price: price,
	}

	if err := db.Save(&product).Error; err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("product:%d", id)
	client.HMSet(c, cacheKey, map[string]any{
		"name":  name,
		"price": price,
	})

	client.Expire(c, cacheKey, time.Minute)

	return nil
}

func invalidDateProductCache(id uint) error {
	cacheKey := fmt.Sprintf("product:%d", id)
	_, err := client.Del(c, cacheKey).Result()
	return err
}

func deleteProductEventBased(db *gorm.DB, id uint) error {
	if err := db.Delete(&Product{}, id).Error; err != nil {
		return err
	}

	return invalidDateProductCache(id)
}

func getRecentProducts(db *gorm.DB) ([]Product, error) {
	productIds, err := client.LRange(c, "recent_products", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var products []Product
	for _, idStr := range productIds {
		id, _ := strconv.Atoi(idStr)
		product, err := getProductByIdHash(db, uint(id))
		if err == nil {
			products = append(products, product)
		}
	}

	return products, nil
}

func updateProductWithTransaction(db *gorm.DB, id uint, name string, price int) error {
	cacheKey := fmt.Sprintf("product:%d", id)

	_, err := client.TxPipelined(c, func(pipe redis.Pipeliner) error {
		pipe.HMSet(c, cacheKey, map[string]any{
			"name":  name,
			"price": price,
		})
		pipe.Expire(c, cacheKey, time.Minute)

		return nil
	})

	if err == nil {
		db.Model(&Product{}).Where("id = ?").Updates(Product{
			Name:  name,
			Price: price,
		})
	}

	return err
}

func RedisDb() {
	db := initDB()
	r := gin.Default()

	r.POST("/product", func(c *gin.Context) {
		var req struct {
			ID    uint   `json:"id"`
			Name  string `json:"name"`
			Price int    `json:"price"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		err := createOrUpdateProductWriteThrough(db, req.ID, req.Name, req.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "created/updated"})
	})

	r.DELETE("/product/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := deleteProductEventBased(db, uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	})

	r.POST("/product/invalidate/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := invalidDateProductCache(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "cache invalidated"})
	})

	r.GET("/product/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		product, err := getProductByIdHash(db, uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		addToRecentProductList(uint(id))
		c.JSON(http.StatusOK, product)
	})

	r.PUT("/product/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var req struct {
			Name  string `json:"name"`
			Price int    `json:"price"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if err := updateProductWithTransaction(db, uint(id), req.Name, req.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "updated"})
	})

	r.GET("/recent_products", func(c *gin.Context) {
		products, err := getRecentProducts(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	r.Run(":8080")
}
