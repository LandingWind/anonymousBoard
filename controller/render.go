package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	if hash != "" {
		c.HTML(http.StatusOK, "content.html", gin.H{
			"hash": hash,
		})
	} else {
		c.HTML(http.StatusOK, "404.html", gin.H{})
	}
}
