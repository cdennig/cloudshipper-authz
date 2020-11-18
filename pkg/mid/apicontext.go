package mid

import (
	apicontext "github.com/cdennig/cloudshipper-authz/pkg/context"
	"github.com/gin-gonic/gin"
)

// APIContext - middleware for API context (from header values)
func APIContext() gin.HandlerFunc {
	return func(c *gin.Context) {

		var apiCtx apicontext.APIContext

		if err := c.ShouldBindHeader(&apiCtx); err != nil {
			c.AbortWithStatusJSON(403, gin.H{"error": "Either tenant or user haven't been set."})
			return
		}
		// Set shared variable between handlers
		c.Set("csTenant", apiCtx.Tenant)
		c.Set("csUser", apiCtx.User)

		c.Next()
	}
}
