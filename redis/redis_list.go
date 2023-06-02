package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func RedisToList(ctx context.Context, client *redis.Client) {
	key := "ids"
	value := []interface{}{1, 2, 3, "Kol"}
	err := client.RPush(ctx, key, value...).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0,
	})
	RedisToList(context.Background(), client)
}
