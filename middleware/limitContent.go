package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	constant "github.com/wkk5194/anonymousBoard/constant"
)

func LimitContentLength(c *gin.Context) {
	content := c.PostForm("content")
	if (len(content) >= constant.LimitContentLength) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "content too long",
		})
		return
	}
	c.Next()
}
