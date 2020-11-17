package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// InjectTraceID 注入一个 traceid
func InjectTraceID(c *gin.Context) {
	var traceID string
	if traceID = c.Writer.Header().Get("X-Request-Id"); traceID == "" { // 为了形成统一链路
		traceID = uuid.NewV4().String()
		c.Writer.Header().Set("X-Request-Id", traceID)
	}
	fmt.Printf("trace with %s\n", traceID)
	c.Next()
}
