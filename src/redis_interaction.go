package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func push_data(key string, value interface{}) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		log.Panic(err)
	}
	pong, err := client.Ping().Result()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(pong)

	err = client.Set(key, cacheEntry, 0).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
	}
}
