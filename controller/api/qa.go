package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/wkk5194/anonymousBoard/model"
)

func GetQA(c *gin.Context) {
	qa := model.GetQA()
	if qa == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"qa":      qa,
	})
}
