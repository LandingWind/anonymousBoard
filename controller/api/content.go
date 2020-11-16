package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/wkk5194/anonymousBoard/model"
)

func CreateContent(c *gin.Context) {
	// key := c.PostForm("key")
	hash := model.CreateContent(nil)
	c.JSON(http.StatusOK, gin.H{
		"contentUrl": fmt.Sprintf("http://localhost:9001/content/%s", hash),
	})
}
