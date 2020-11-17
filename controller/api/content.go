package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/wkk5194/anonymousBoard/model"
)

func CreateContent(c *gin.Context) {
	// TODO: change to bind
	// post data
	masterKey := c.PostForm("masterKey")
	content := c.PostForm("content")
	lock := c.PostForm("lock")
	editable := c.PostForm("editable")
	fmt.Printf("%s %s %s %s\n", masterKey, content, lock, editable)
	//
	hash, err := model.CreateContent(masterKey, content, lock, editable)
	fmt.Printf("hash: %s\n", hash)
	if err != nil {
		c.HTML(http.StatusOK, "404.html", gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"contentUrl": fmt.Sprintf("http://localhost:9001/content/%s", hash),
	})
}

func SaveContent(c *gin.Context) {
	content := c.PostForm("content")
	hash := c.PostForm("hash")
	if !model.QueryHMKeyExist(hash) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "no this message token",
		})
		return
	}
	rds := model.GetRedis()
	err := rds.HSet(hash, "content", content).Err()
	fmt.Printf("%j\n", err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "save error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func GetContent(c *gin.Context) {
	hash := c.PostForm("hash")
	content := model.GetHMKeyValue(hash)
	if content == nil || len(content) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "no this message token",
		})
		return
	}
	if content["lock"] == "false" {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"content": content["content"],
		})
		return
	}
	masterKey := c.PostForm("masterKey")
	if content["masterKey"] == masterKey {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"content": content["content"],
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"msg":     "wrong masterKey",
	})
}
