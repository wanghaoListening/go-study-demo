package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

//go get -u github.com/go-redis/redis

var redisClient *redis.Client

func ConnRedis() {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping().Result()

	if err != nil {
		fmt.Println("connect redis failed")
		return
	}
}

func Get(key string) string {

	result, err := redisClient.Get(key).Result()
	if err != nil {
		fmt.Println("Get failed")
	}
	return result
}
