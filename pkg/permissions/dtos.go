package permissions

// CheckPermissionsResultItemDTO - Result item result
type CheckPermissionsResultItemDTO struct {
	Allowed bool `json:"allowed"`
}

// CheckPermissionsItemDTO - Item to check permissions
type CheckPermissionsItemDTO struct {
	Domain string `json:"domain" binding:"required"`
	Type   string `json:"type" binding:"required"`
	ID     string `json:"id" binding:"required"`
	Action string `json:"action" binding:"required,oneof=read write *"`
}

// CheckPermissionsResultDTO - List of result items
type CheckPermissionsResultDTO struct {
	Items []CheckPermissionsResultItemDTO `json:"items"`
}

// CheckPermissionsDTO - List of result items
type CheckPermissionsDTO struct {
	Items []CheckPermissionsItemDTO `json:"items" binding:"required,max=10,min=1,dive"`
}

// ImplicitPermissionsResultDTO - result of implicit handler
type ImplicitPermissionsResultDTO struct {
	ImplicitRoles       []string   `json:"implicitRoles"`
	ImplicitPermissions [][]string `json:"implicitPermissions"`
}

// ImplicitPermissionsItemDTO - query params
type ImplicitPermissionsItemDTO struct {
	Domain string `url:"domain" binding:"required"`
}

// AssignmentItemDTO - query params
type AssignmentItemDTO struct {
	Domain string `url:"domain" binding:"required"`
}

// AssignmentResultDTO - result of implicit handler
type AssignmentResultDTO struct {
	Assignments [][]string `json:"assignments"`
}
