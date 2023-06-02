package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	//conn, err := redis.Dial("tcp", "0.0.0.0:6379")
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//fmt.Println("redis 链接成功")
	//defer conn.Close()

	//获取redis链接
	client := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0,
	})
	//val, err := client.Get(context.Background(), "a").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("a:", val)
	//new_val, err := client.Get(context.Background(), "ruling").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(new_val)
	err := client.Set(context.Background(), "go_redis_test_key", "go_redis_test_value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get(context.Background(), "go_redis_test_key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("go_redis_test_key: ", val)
	do_val, err := client.Do(context.Background(), "get", "go_redis_test_key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("do  go_redis_test_key: ", do_val)
}
