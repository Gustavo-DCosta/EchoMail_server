package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client // <-- EXPORT this
	Ctx = context.Background()
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error opening .env file", err)
	}
}

func Connect_To_Redis() {
	Red_Addr := os.Getenv("Redis_Adr")
	Red_Pass := os.Getenv("Redis_Pass")

	Rdb = redis.NewClient(&redis.Options{
		Addr:     Red_Addr,
		Username: "default",
		Password: Red_Pass,
		DB:       0,
	})

	// Quick check
	err := Rdb.Set(Ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	result, err := Rdb.Get(Ctx, "foo").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(result) // >>> bar
	fmt.Println("Connected to redis")
}
