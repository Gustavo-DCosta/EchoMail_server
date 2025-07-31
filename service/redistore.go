package service

import (
	"fmt"
	"time"

	"github.com/Gustavo-DCosta/server/database"
	"github.com/redis/go-redis/v9"
)

func WriteRedis2ways(uuid, phoneNumber string) error {
	err := database.Rdb.Set(database.Ctx, phoneNumber, uuid, 10*time.Minute).Err()
	if err != nil {
		fmt.Println("Error writting to redis", err)
		return err
	}
	err = database.Rdb.Set(database.Ctx, uuid, phoneNumber, 10*time.Minute).Err()
	if err != nil {
		fmt.Println("Error writting to redis", err)
		return err
	}

	return nil
}

func CrossCheck(RedisPhoneNumber string) (string, error) {
	val, err := database.Rdb.Get(database.Ctx, RedisPhoneNumber).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist", RedisPhoneNumber)
	} else if err != nil {
		// Return the actual error instead of using log.Fatal
		return "", fmt.Errorf("redis error: %w", err)
	}

	return val, nil
}
