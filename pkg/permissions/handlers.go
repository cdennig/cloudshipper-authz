package permissions

import (
	"github.com/casbin/casbin/v2"
	"github.com/cdennig/cloudshipper-authz/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PermissionHandler struct {
	enforcer *casbin.CachedEnforcer
}

func NewPermissionHandler(enforcer *casbin.CachedEnforcer) *PermissionHandler {
	return &PermissionHandler{enforcer: enforcer}
}

// Check - check permissions
func (p *PermissionHandler) Check(c *gin.Context) {
	tenant := c.Keys["csTenant"].(string)
	user := c.Keys["csUser"].(string)
	var itemsToCheck CheckPermissionsDTO
	err := c.ShouldBindJSON(&itemsToCheck)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors := validation.WrapAPIValidationErrors(errs)

			c.AbortWithStatusJSON(400, validationErrors)
			return
		}

		c.AbortWithError(500, err)
		return
	}

	var result CheckPermissionsResultDTO

	for _, v := range itemsToCheck.Items {
		domainToCheck := tenant + "/" + v.Domain
		ok, err := p.enforcer.Enforce(user, domainToCheck, v.Type, v.ID, v.Action)

		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		if ok == true {
			var res CheckPermissionsResultItemDTO
			res.Allowed = true
			result.Items = append(result.Items, res)
		} else {
			var res CheckPermissionsResultItemDTO
			res.Allowed = false
			result.Items = append(result.Items, res)
		}
	}

	c.JSON(200, result)

}

// // Implicit - Get implicit permissions & roles
// func (p *PermissionHandler) Implicit(c *gin.Context) {
// 	tenant := ctx.Values().GetString("csTenant")
// 	user := ctx.Values().GetString("csUser")
// 	var input ImplicitPermissionsItemDTO
// 	err := ctx.ReadQuery(&input)
// 	if err != nil {

// 		if errs, ok := err.(validator.ValidationErrors); ok {
// 			validationErrors := validation.WrapAPIValidationErrors(errs)

// 			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
// 				Title("Validation error").
// 				Detail("One or more fields failed to be validated").
// 				Key("errors", validationErrors))

// 			return
// 		}

// 		ctx.StopWithStatus(iris.StatusInternalServerError)
// 		return
// 	}

// 	if input.Domain == "" {
// 		input.Domain = "*"
// 	}

// 	var result ImplicitPermissionsResultDTO
// 	impRoles, _ := p.enforcer.GetImplicitRolesForUser(user, tenant+"/"+input.Domain)
// 	impPermissions, _ := p.enforcer.GetImplicitPermissionsForUser(user, tenant+"/"+input.Domain)

// 	result.ImplicitRoles = impRoles
// 	result.ImplicitPermissions = impPermissions

// 	ctx.StatusCode(200)
// 	ctx.JSON(result)
// }

// // Assignment - Get assignments
// func (p *PermissionHandler) Assignment(c *gin.Context) {
// 	user := ctx.Values().GetString("csUser")
// 	var input ImplicitPermissionsItemDTO
// 	err := ctx.ReadQuery(&input)
// 	if err != nil {

// 		if errs, ok := err.(validator.ValidationErrors); ok {
// 			validationErrors := validation.WrapAPIValidationErrors(errs)

// 			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
// 				Title("Validation error").
// 				Detail("One or more fields failed to be validated").
// 				Key("errors", validationErrors))

// 			return
// 		}

// 		ctx.StopWithStatus(iris.StatusInternalServerError)
// 		return
// 	}

// 	if input.Domain == "" {
// 		input.Domain = "*"
// 	}

// 	var result AssignmentResultDTO
// 	filteredGroupingPolicy := p.enforcer.GetFilteredGroupingPolicy(0, user)

// 	result.Assignments = filteredGroupingPolicy

// 	ctx.StatusCode(200)
// 	ctx.JSON(result)
// }
