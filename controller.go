package anonymousboard

import "github.com/gin-gonic/gin"

// ShowIndex 显示主页的渲染
func ShowIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
