package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/wkk5194/anonymousBoard/model"
)

// ShowIndex 显示主页的渲染
func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"title": "Main website",
	})
}

// ShowContent 显示内容
func ShowContent(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"msg": "not support empty message token",
		})
		return
	}
	content := model.GetContent(hash)
	if len(content) == 0 {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"msg": "not have this message",
		})
		return
	}
	c.HTML(http.StatusOK, "content.html", gin.H{
		"hash":     hash,
		"content":  content["content"],
		"editable": content["editable"] == "true",
		"lock":     content["lock"] == "true",
	})
}
