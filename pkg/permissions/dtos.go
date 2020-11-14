package permissions

// CheckPermissionsResultItemDTO - Result item result
type CheckPermissionsResultItemDTO struct {
	Allowed bool `json:"allowed"`
}

// CheckPermissionsItemDTO - Item to check permissions
type CheckPermissionsItemDTO struct {
	Domain string `json:"domain" validate:"required"`
	Type   string `json:"type" validate:"required"`
	ID     string `json:"id" validate:"required"`
	Action string `json:"action" validate:"required,oneof=read write *"`
}

// CheckPermissionsResultDTO - List of result items
type CheckPermissionsResultDTO struct {
	Items []CheckPermissionsResultItemDTO `json:"items"`
}

// CheckPermissionsDTO - List of result items
type CheckPermissionsDTO struct {
	Items []CheckPermissionsItemDTO `json:"items" validate:"required,max=10,min=1,dive"`
}

// ImplicitPermissionsResultDTO - result of implicit handler
type ImplicitPermissionsResultDTO struct {
	ImplicitRoles       []string   `json:"implicitRoles"`
	ImplicitPermissions [][]string `json:"implicitPermissions"`
}

// ImplicitPermissionsItemDTO - query params
type ImplicitPermissionsItemDTO struct {
	Domain string `url:"domain" validate:"required"`
}

// AssignmentItemDTO - query params
type AssignmentItemDTO struct {
	Domain string `url:"domain" validate:"required"`
}

// AssignmentResultDTO - result of implicit handler
type AssignmentResultDTO struct {
	Assignments [][]string `json:"assignments"`
}
