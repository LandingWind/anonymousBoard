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
	flag := QueryKeyExist(key)
	if flag {
		val, err := rdb.Get(key).Result()
		if err != nil {
			panic(err)
		}
		return val
	}
	return nil
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
