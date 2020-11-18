package server

import (
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/cdennig/cloudshipper-authz/internal/healthz"
	prom "github.com/cdennig/cloudshipper-authz/internal/prometheus"
	"github.com/cdennig/cloudshipper-authz/pkg/permissions"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Routing setup api routing
func Routing(enforcer *casbin.CachedEnforcer /*db *DB,*/) *gin.Engine {

	r := gin.Default()
	healthz.RegisterHandlers(r)
	prom.RegisterHandlers(r)
	permissions.RegisterHandlers(enforcer, r)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "DELETE", "PUT", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "accept", "origin", "Cache-Control", "X-Cs-Tenant", "X-Cs-User"},
		ExposeHeaders:    []string{"Content-Length", "Location"},
		AllowCredentials: true,
		MaxAge:           15 * time.Minute,
	}))
	return r
}
