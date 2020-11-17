package model

import (
	"fmt"

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
func CreateContent(masterKey string, content string, lock string, editable string) (string, error) {
	newHash := GetNewHash()
	// store to redis
	storeValue := make(map[string]interface{})
	storeValue["masterKey"] = masterKey
	storeValue["content"] = content
	storeValue["lock"] = lock
	storeValue["editable"] = editable
	rdb := GetRedis()
	setErr := rdb.HMSet(newHash, storeValue).Err()
	if setErr != nil {
		fmt.Println("insert into redis error")
		return "", setErr
	}
	return newHash, nil
}
