package main

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/wkk5194/anonymousBoard/controller"
)

func main() {
	r := gin.Default()
	r.GET("/welcome", ctrl.ShowIndex)

	r.Run("http://localhost:9001")
}
