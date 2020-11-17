package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	defaultMsg     string = "Something Wrong Occurred!"
	defaultErrCode int    = 400
	defaultTitle   string = "Error"
)

// RenderErrorPage 渲染错误页面
func RenderErrorPage(c *gin.Context) {
	c.Next()
	// check errorPage
	if tval, ok := c.Get("errorPage"); ok {
		val := tval.(gin.H)
		var title string
		var code int
		var msg string
		if _, ok := val["title"]; !ok {
			title = defaultTitle
		} else {
			title = val["title"].(string)
		}
		if _, ok := val["code"]; !ok {
			code = defaultErrCode
		} else {
			code = val["code"].(int)
		}
		if _, ok := val["msg"]; !ok {
			msg = defaultMsg
		} else {
			msg = val["msg"].(string)
		}
		c.HTML(http.StatusOK, "error.html", gin.H{
			"msg":   msg,
			"code":  fmt.Sprintf("%d", code),
			"title": title,
		})
	}
	// check other error
	if c.Writer.Status() != 200 {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"msg":   defaultMsg,
			"code":  fmt.Sprintf("%d", defaultErrCode),
			"title": defaultTitle,
		})
	}
}
