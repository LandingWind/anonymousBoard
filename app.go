package main

import (
	"time"

	"github.com/gin-gonic/gin"
	ctrl "github.com/wkk5194/anonymousBoard/controller"
	api "github.com/wkk5194/anonymousBoard/controller/api"
	middleware "github.com/wkk5194/anonymousBoard/middleware"
	model "github.com/wkk5194/anonymousBoard/model"
	"golang.org/x/time/rate"
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
	// before handle main logic
	r.Use(middleware.InjectTraceID)
	// after handle main logic
	r.Use(middleware.RenderErrorPage)

	// render router
	// begin page
	r.GET("/", ctrl.ShowIndex)
	r.GET("/welcome", ctrl.ShowIndex)
	// content page
	r.GET("/content/:hash", ctrl.ShowContent)

	// api router
	contentAPI := r.Group("/api/content")
	contentAPI.Use(middleware.LimitContentLength)
	contentAPI.Use(middleware.RateLimiter(
		func(c *gin.Context) string {
			return c.ClientIP() // limit rate by client ip
		}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
			// limit 1 qps/clientIp and permit bursts of at most 10 tokens
			// and the limiter liveness time duration is 1 hour
			return rate.NewLimiter(rate.Every(1000*time.Millisecond), 10), time.Hour
		},
	))
	{
		contentAPI.POST("/create", api.CreateContent)
		contentAPI.POST("/save", api.SaveContent)
		contentAPI.POST("/get", api.GetContent)
	}
	r.GET("/api/qa", api.GetQA)

	r.Run(":9001")
}
