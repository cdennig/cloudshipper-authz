package mid

import (
	"github.com/gin-gonic/gin"
)

type versionHeader struct {
	Version string `header:"X-Cs-Version" binding:"required,oneof='v1'"`
}

// APIVersion - middleware for API version (from header values)
func APIVersion() gin.HandlerFunc {
	return func(c *gin.Context) {

		var versionCtx versionHeader
		if err := c.ShouldBindHeader(&versionCtx); err != nil {
			c.JSON(400, gin.H{"error": "API version header not set."})
			return
		}

		c.Set("csVersion", versionCtx.Version)
		c.Next()

	}
}
