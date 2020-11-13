package mid

import (
	"net/http"

	apicontext "github.com/cdennig/cloudshipper-authz/pkg/context"
	"github.com/kataras/iris/v12"
)

// APIContext - middleware for API context (from header values)
func APIContext() iris.Handler {
	return func(ctx iris.Context) {

		var apiCtx apicontext.APIContext
		if err := ctx.ReadHeaders(&apiCtx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
		// Set shared variable between handlers
		ctx.Values().Set("csTenant", apiCtx.Tenant)
		ctx.Values().Set("csUser", apiCtx.User)

		ctx.Next()
	}
}
