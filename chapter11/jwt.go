package chapter11

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

var jwtKey = []byte("secret_key")

func JwtGin() {
	r := gin.Default()
	defer r.Run()

	r.POST("/login", loginHandle)

	auth := r.Group("/auth")
	auth.Use(authmiddleware())
	{
		auth.GET("/hello", helloHandle)
	}
}

func loginHandle(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loginData.Username != "skarekroe" || loginData.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid login credentials"})
		return
	}

	expiryTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: loginData.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func helloHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, authenticated user"})
}

func authmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}
