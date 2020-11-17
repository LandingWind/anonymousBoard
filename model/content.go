package model

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
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
	storeValue["masterKey"] = masterKey                                    // 访问密钥
	storeValue["content"] = content                                        // 内容
	storeValue["lock"] = lock                                              // 是否上锁
	storeValue["expireTime"] = time.Now().Add(time.Hour * 24 * 7).String() // 失效时间
	storeValue["visitTime"] = 0                                            // 访问次数
	storeValue["lastEditTime"] = time.Now().String()                       // 最后一次编辑的时间
	rdb := GetRedis()
	setErr := rdb.HMSet(newHash, storeValue).Err()
	if setErr != nil {
		fmt.Println("insert into redis error")
		return "", setErr
	}
	return newHash, nil
}
