package prom

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// RegisterHandlers register handler for health check endpoint
func RegisterHandlers(router *gin.Engine) {
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)
}
