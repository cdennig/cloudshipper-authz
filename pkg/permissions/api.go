package permissions

import (
	"github.com/casbin/casbin/v2"
	"github.com/cdennig/cloudshipper-authz/pkg/mid"
	"github.com/kataras/iris/v12"
)

// RegisterHandlers register handler for permissions endpoints
func RegisterHandlers(enforcer *casbin.CachedEnforcer, app *iris.Application) {

	permissionAPI := app.Party("/permissions")
	{
		permissionHandler := NewPermissionHandler(enforcer)
		permissionAPI.Use(mid.APIContext())
		permissionAPI.Use(mid.APIVersion())
		permissionAPI.Post("/check", permissionHandler.Check)
	}
}
