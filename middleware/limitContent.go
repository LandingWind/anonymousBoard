package middleware

import (
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	constant "github.com/wkk5194/anonymousBoard/constant"
)

func LimitContentLength(c *gin.Context) {
	content := c.PostForm("content")
	len := utf8.RuneCountInString(content)
	if (len >= constant.LimitContentLen) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "content too long",
		})
		c.Abort()
		return
	}
	c.Next()
}
