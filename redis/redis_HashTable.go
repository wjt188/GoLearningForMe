package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func RedisToHashTable(ctx context.Context, client *redis.Client) {
	//err := client.HSet(ctx, "学生1", "Name", "Curry", "Height", 191).Err()
	//if err != nil {
	//	panic(err)
	//}
	err := client.HSet(ctx, "学生2", "Name", "James", "Height", 206).Err()
	if err != nil {
		panic(err)
	}
	//for field, value := range client.HGetAll(ctx, "学生1").Val() {
	//	fmt.Println(field, ":", value)
	//}

	for field, value := range client.HGetAll(ctx, "学生2").Val() {
		fmt.Println(field, ":", value)
	}
	client.Del(ctx, "学生1")

}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0,
	})
	RedisToHashTable(context.Background(), client)
}
