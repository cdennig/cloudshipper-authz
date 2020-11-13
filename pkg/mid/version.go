package mid

import (
	"github.com/cdennig/cloudshipper-authz/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type versionHeader struct {
	Version string `header:"X-Cs-Version" validate:"required,oneof='v1'"`
}

// APIVersion - middleware for API version (from header values)
func APIVersion() iris.Handler {
	return func(ctx iris.Context) {

		var versionCtx versionHeader
		if err := ctx.ReadHeaders(&versionCtx); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				validationErrors := validation.WrapAPIValidationErrors(errs)
				ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
					Title("Validation error").
					Detail("One or more fields failed to be validated").
					Key("errors", validationErrors))
				return
			}

		} else {
			ctx.Values().Set("csVersion", versionCtx.Version)
			ctx.Next()
		}
	}
}
