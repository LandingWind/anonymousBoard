package model

import (
	uuid "github.com/satori/go.uuid"
)

func GetContent(hash string) interface{} {
	// use redis to fetch content
	val := GetKeyValue(hash)
	return val
}

// GetNewHash 获取一个新的 hash string
func GetNewHash() string {
	// hash length is 8
	fixedLen := 8
	for {
		u1 := uuid.NewV4().String()
		u2 := u1[0:fixedLen]
		if !QueryKeyExist(u2) {
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
func CreateContent(keys map[string]string) string {
	newHash := GetNewHash()
	// store to redis
	// storeValue := make(map[string]string)
	SetKeyValue(newHash, "abc") // TODO: imple
	return newHash
}
