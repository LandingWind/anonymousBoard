package middleware

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

// 每5分钟过期，每10分钟清理
var limiterSet = cache.New(5*time.Minute, 10*time.Minute)

// 基于某一个key限制访问频率
func RateLimiter(
	key func(*gin.Context) string,
	createLimiter func(*gin.Context) (*rate.Limiter, time.Duration),
) gin.HandlerFunc {
	return func(c *gin.Context) {
		k := key(c) // rate key
		limiter, ok := limiterSet.Get(k)
		if !ok {
			var expire time.Duration
			limiter, expire = createLimiter(c)
			limiterSet.Set(k, limiter, expire)
		}
		ok = limiter.(*rate.Limiter).Allow() // check 是否可以通行
		if !ok {
			c.HTML(http.StatusOK, "error.html", gin.H{
				"msg":   "too many requests at the same time",
				"code":  500,
				"title": defaultTitle,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
