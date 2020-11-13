package server

import (
	"github.com/casbin/casbin/v2"
	"github.com/cdennig/cloudshipper-authz/internal/healthz"
	prom "github.com/cdennig/cloudshipper-authz/internal/prometheus"
	"github.com/cdennig/cloudshipper-authz/pkg/permissions"
	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/cors"
	prometheusMiddleware "github.com/iris-contrib/middleware/prometheus"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Routing setup api routing
func Routing(enforcer *casbin.CachedEnforcer /*db *DB,*/) *iris.Application {
	app := iris.New()
	m := prometheusMiddleware.New("cs-authz", 0.05, 0.2, 0.5, 1, 3)
	app.Use(m.ServeHTTP)
	app.Validator = validator.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(iris.Compression)
	app.AllowMethods(iris.MethodOptions)
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "DELETE", "PUT", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Content-Length", "Location"},
		MaxAge:           600,
	})
	app.Use(crs)

	healthz.RegisterHandlers(app)
	prom.RegisterHandlers(app)
	permissions.RegisterHandlers(enforcer, app)
	return app
}
