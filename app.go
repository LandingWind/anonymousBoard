package main

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/wkk5194/anonymousBoard/controller"
)

func main() {
	r := gin.Default()
	// load views
	r.LoadHTMLGlob("views/*")

	// router
	// begin page
	r.GET("/", ctrl.ShowIndex)
	r.GET("/welcome", ctrl.ShowIndex)
	// content page
	r.GET("/content/:hash", ctrl.ShowContent)

	r.Run(":9001")
}
