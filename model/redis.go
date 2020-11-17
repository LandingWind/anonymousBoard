package model

import (
	"fmt"

	"github.com/go-redis/redis"
	constant "github.com/wkk5194/anonymousBoard/constant"
)

// RedisClient redis连接池
var rdb *redis.Client

// InitRedis redis初始化
func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", constant.REDIS_HOST, constant.REDIS_PORT),
		Password: constant.REDIS_AUTH,
		DB:       0,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
}

// QueryHMKeyExist 查询是否存在某一个key for hashmap
func QueryHMKeyExist(key string) bool {
	val, err := rdb.HGetAll(key).Result()
	if err == redis.Nil || len(val) == 0 {
		return false
	} else if err != nil {
		panic(err)
	}
	return true
}

// GetHMKeyValue 获取value for hashmap
func GetHMKeyValue(key string) map[string]string {
	val, err := rdb.HGetAll(key).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		panic(err)
	}
	return val
}

// SetHMKeyValue 设置key-value for hashmap
func SetHMKeyValue(key string, value map[string]interface{}) bool {
	err := rdb.HMSet(key, value).Err()
	if err != nil {
		fmt.Printf("set key-value for hashmap error\n")
		return false
	}
	return true
}

// QueryKeyExist 查询是否存在某一个key
func QueryKeyExist(key string) bool {
	_, err := rdb.Get(key).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		panic(err)
	}
	return true
}

// GetKeyValue 获取value
func GetKeyValue(key string) interface{} {
	val, err := rdb.Get(key).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		panic(err)
	}
	return val
}

// SetKeyValue 设置key-value
func SetKeyValue(key string, value interface{}) bool {
	err := rdb.Set(key, value, 0).Err()
	if err != nil {
		fmt.Printf("set key-value:%s-%s error\n", key, value)
		return false
	}
	return true
}

// GetRedis 直接获取 redis连 接进行操作
func GetRedis() *redis.Client {
	return rdb
}
