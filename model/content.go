package model

import (
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	timeFormat string = "2006-01-02 15:04:05"
)

// GetContent 获取数据
func GetContent(hash string) map[string]string {
	// use redis to fetch content
	val := GetHMKeyValue(hash)
	return val
}

// GetNewHash 获取一个新的 hash string
func GetNewHash() string {
	// hash length is 8
	fixedLen := 8
	for {
		u1 := uuid.NewV4().String()
		u2 := u1[0:fixedLen]
		if !QueryHMKeyExist(u2) {
			return u2
		}
	}
	panic("not exist free hash string")
}

/*
CreateContent will return new hash string
accept params:
enterKey(optional)
editKey(optional)
masterKey
*/
func CreateContent(masterKey string, content string, lock string) (string, error) {
	newHash := GetNewHash()
	// store to redis
	storeValue := make(map[string]interface{})
	storeValue["masterKey"] = masterKey                                              // 访问密钥
	storeValue["content"] = content                                                  // 内容
	storeValue["lock"] = lock                                                        // 是否上锁
	storeValue["expireTime"] = time.Now().Add(time.Hour * 24 * 7).Format(timeFormat) // 失效时间
	storeValue["visitTime"] = 0                                                      // 访问次数
	storeValue["lastEditTime"] = time.Now().Format(timeFormat)                       // 最后一次编辑的时间
	rdb := GetRedis()
	dbSize, err := rdb.DBSize().Result()
	if (err != nil || dbSize >= 10000) {
		return "", errors.New("redis db full, please try later")
	}
	setErr := rdb.HMSet(newHash, storeValue).Err()
	if setErr != nil {
		fmt.Println("insert into redis error")
		return "", setErr
	}
	return newHash, nil
}

// UpdateLastEditTime 更新最后一次更新时间
func UpdateLastEditTime(hash string) {
	time := time.Now().Format(timeFormat)
	GetRedis().HSet(hash, "lastEditTime", time)
}

// UpdateVisitTime 更新 visitTime
func UpdateVisitTime(hash string) {
	GetRedis().HIncrBy(hash, "visitTime", 1)
}
