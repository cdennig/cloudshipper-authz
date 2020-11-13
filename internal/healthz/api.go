package healthz

import (
	"github.com/kataras/iris/v12"
)

// RegisterHandlers register handler for health check endpoint
func RegisterHandlers(app *iris.Application) {
	healthzAPI := app.Party("/healthz")
	{
		healthzAPI.Get("/", check)
	}
}

// Check implements endpoint for health check
func check(ctx iris.Context) {
	ctx.StatusCode(200)
}
