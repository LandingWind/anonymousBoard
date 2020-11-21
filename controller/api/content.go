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
	//
	hash, err := model.CreateContent(masterKey, content, lock)
	fmt.Printf("hash: %s\n", hash)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"contentUrl": fmt.Sprintf("/content/%s", hash),
	})
}

func SaveContent(c *gin.Context) {
	contentData := c.PostForm("content")
	hash := c.PostForm("hash")
	content := model.GetHMKeyValue(hash)
	if content == nil || len(content) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "no this message token",
		})
		return
	}
	if content["lock"] == "true" && content["masterKey"] != c.PostForm("masterKey") {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "wrong masterKey",
		})
		return
	}
	rds := model.GetRedis()
	err := rds.HSet(hash, "content", contentData).Err()
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
	// update lastEdit time
	model.UpdateLastEditTime(hash)
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
		// update visit time
		model.UpdateVisitTime(hash)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"msg":     "wrong masterKey",
	})
}
