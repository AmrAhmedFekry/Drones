package Cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

var ctx = context.Background()

// Connect to redis
func ConnectRedis() *Redis {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	redis := Redis{
		client: rdb,
	}
	return &redis
}

// Set Specific key and value to redis
func SetToRedis(key string, value interface{}, ttl int) {
	cachedValue, _ := json.Marshal(value)
	// Set the key with expiry time
	errInSet := ConnectRedis().client.Set(ctx, key, cachedValue, time.Duration(ttl)*time.Second).Err()
	if errInSet != nil {
		panic(errInSet)
	}
}

// Get Specific key from redis
func GetFromRedis(key string) interface{} {
	val, err := ConnectRedis().client.Get(ctx, key).Result()
	if err != nil {
		panic(err.Error())
	}
	var result interface{}
	json.Unmarshal([]byte(val), &result)
	return result
}
