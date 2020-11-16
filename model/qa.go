package model

func GetQA() interface{} {
	qa := GetKeyValue("qa")
	return qa
}
