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

	/*
		-> Store in the reverse way so later we can validate the propierty of the jwt by doing coss checking
	*/

	err = database.Rdb.Set(database.Ctx, uuid, phoneNumber, 10*time.Minute).Err()
	if err != nil {
		fmt.Println("Error writting to redis", err)
		return err
	}

	return nil
}
func CrossUuidToPhone(redisUuid string) (string, error) {
	val, err := database.Rdb.Get(database.Ctx, redisUuid).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist", redisUuid)
	} else if err != nil {
		// Return the actual error instead of using log.Fatal
		return "", fmt.Errorf("redis error: %w", err)
	}

	return val, nil
}
func CrossPhonetoUuid(redisPhoneNumber string) (string, error) {
	val, err := database.Rdb.Get(database.Ctx, redisPhoneNumber).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist", redisPhoneNumber)
	} else if err != nil {
		// Return the actual error instead of using log.Fatal
		return "", fmt.Errorf("redis error: %w", err)
	}

	return val, nil
}
