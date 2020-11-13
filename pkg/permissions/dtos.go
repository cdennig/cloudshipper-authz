package permissions

// CheckPermissionsResultItemDTO - Result item result
type CheckPermissionsResultItemDTO struct {
	Allowed bool `json:"allowed"`
	// ImpRoles            []string
	// AllRoles            []string
	// Roles               []string
	// Permissions         [][]string
	// ImpPermissions      [][]string
	// NamedGroupingPolicy [][]string
	// FilteredGroupingPolicy [][]string
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
