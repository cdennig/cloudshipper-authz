package permissions

import (
	"github.com/casbin/casbin/v2"
	"github.com/cdennig/cloudshipper-authz/pkg/mid"
	"github.com/gin-gonic/gin"
)

// RegisterHandlers register handler for permissions endpoints
func RegisterHandlers(enforcer *casbin.CachedEnforcer, router *gin.Engine) {

	permissionAPI := router.Group("/permissions")
	{
		permissionHandler := NewPermissionHandler(enforcer)
		permissionAPI.Use(mid.APIContext())
		permissionAPI.Use(mid.APIVersion())
		permissionAPI.POST("/check", permissionHandler.Check)
		// permissionAPI.GET("/implicit", permissionHandler.Implicit)
		// permissionAPI.GET("/assignments", permissionHandler.Assignment)
	}
}
