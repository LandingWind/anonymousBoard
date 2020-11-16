package anonymousboard

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/welcome", ShowIndex)

	r.Run("http://localhost:9001")
}
