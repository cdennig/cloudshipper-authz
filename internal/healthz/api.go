package healthz

import (
	"github.com/gin-gonic/gin"
)

// RegisterHandlers register handler for health check endpoint
func RegisterHandlers(router *gin.Engine) {
	healthzAPI := router.Group("/healthz")
	{
		healthzAPI.GET("/", check)
	}
}

// Check implements endpoint for health check
func check(c *gin.Context) {
	c.Status(200)
}
