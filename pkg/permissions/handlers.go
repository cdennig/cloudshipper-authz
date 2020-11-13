package permissions

import (
	"github.com/casbin/casbin/v2"
	"github.com/cdennig/cloudshipper-authz/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type PermissionHandler struct {
	enforcer *casbin.CachedEnforcer
}

func NewPermissionHandler(enforcer *casbin.CachedEnforcer) *PermissionHandler {
	return &PermissionHandler{enforcer: enforcer}
}

// Check - check permissions
func (p *PermissionHandler) Check(ctx iris.Context) {
	tenant := ctx.Values().GetString("csTenant")
	user := ctx.Values().GetString("csUser")

	var itemsToCheck CheckPermissionsDTO
	err := ctx.ReadJSON(&itemsToCheck)
	if err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors := validation.WrapAPIValidationErrors(errs)

			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
				Title("Validation error").
				Detail("One or more fields failed to be validated").
				Key("errors", validationErrors))

			return
		}

		// It's probably an internal JSON error, let's dont give more info here.
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}

	var result CheckPermissionsResultDTO

	for _, v := range itemsToCheck.Items {
		domainToCheck := tenant + "/" + v.Domain
		ok, err := p.enforcer.Enforce(user, domainToCheck, v.Type, v.ID, v.Action)

		if err != nil {
			ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
				Title("Check permissions failed.").DetailErr(err))
			return
		}
		if ok == true {
			var res CheckPermissionsResultItemDTO
			res.Allowed = true
			// res.AllRoles = allRoles
			// res.Roles = roles
			// res.ImpRoles = impRoles
			// res.Permissions = perm
			// res.ImpPermissions = impPerm
			// res.NamedGroupingPolicy = namedGroupingPolicy
			// res.FilteredGroupingPolicy = filteredGroupingPolicy
			result.Items = append(result.Items, res)
		} else {
			var res CheckPermissionsResultItemDTO
			// res.FilteredGroupingPolicy = filteredGroupingPolicy
			// res.NamedGroupingPolicy = namedGroupingPolicy
			res.Allowed = false
			// res.AllRoles = allRoles
			// res.Roles = roles
			// res.ImpRoles = impRoles
			// res.Permissions = perm
			// res.ImpPermissions = impPerm
			result.Items = append(result.Items, res)
		}
	}

	ctx.StatusCode(200)
	ctx.JSON(result)
	// // gr, _ := p.enforcer.GetRolesForUser("user1", "tenant1/*")
	// perm := p.enforcer.GetPermissionsForUser(user)
	// impPerm, err := p.enforcer.GetImplicitPermissionsForUser(user, tenant)
	// allRoles, _ := p.enforcer.GetRolesForUser(user)
	// roles := p.enforcer.GetRolesForUserInDomain(user, tenant)
	// impRoles, _ := p.enforcer.GetImplicitRolesForUser(user, tenant)
	// namedGroupingPolicy := p.enforcer.GetFilteredNamedGroupingPolicy("g", 0, user)
	// filteredGroupingPolicy := p.enforcer.GetFilteredGroupingPolicy(0, user)
	// // r := p.enforcer.GetAllSubjects()

}
