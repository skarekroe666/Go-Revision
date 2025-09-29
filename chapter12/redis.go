package chapter12

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func RedisEx() {

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(val)

	err := rdb.Set(ctx, "key", "val", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(val)

	val, err := rdb.Do(ctx, "get", "key").Result()
	if err != nil {
		if err == redis.Nil {
			log.Println("Key does not exist")
			return
		}
		log.Fatal(err)
	}
	fmt.Println(val.(string))

	err = rdb.Del(ctx, "key").Err()
	if err != nil {
		log.Fatal(err)
	}
}
