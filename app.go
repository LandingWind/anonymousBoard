package main

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/wkk5194/anonymousBoard/controller"
	api "github.com/wkk5194/anonymousBoard/controller/api"
	middleware "github.com/wkk5194/anonymousBoard/middleware"
	model "github.com/wkk5194/anonymousBoard/model"
)

// init: auto run before main function
func init() {
	// init redis connection
	model.InitRedis()
}

func main() {
	r := gin.Default()
	// load views
	r.LoadHTMLGlob("views/*")

	// use middleware
	r.Use(middleware.InjectTraceID)
	r.Use(middleware.RenderErrorPage)

	// render router
	// begin page
	r.GET("/", ctrl.ShowIndex)
	r.GET("/welcome", ctrl.ShowIndex)
	// content page
	r.GET("/content/:hash", ctrl.ShowContent)

	// api router
	contentAPI := r.Group("/api/content")
	{
		contentAPI.POST("/create", api.CreateContent)
		contentAPI.POST("/save", api.SaveContent)
		contentAPI.POST("/get", api.GetContent)
	}
	r.GET("/api/qa", api.GetQA)

	r.Run(":9001")
}
