package prom

import (
	"github.com/kataras/iris/v12"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterHandlers register handler for health check endpoint
func RegisterHandlers(app *iris.Application) {
	promAPI := app.Party("/metrics")
	{
		promAPI.Get("/", iris.FromStd(promhttp.Handler()))
	}
}
