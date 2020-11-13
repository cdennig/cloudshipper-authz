package apicontext

// APIContext - holds the API context information like tenant id, user etc.
type APIContext struct {
	Tenant string `header:"X-Cs-Tenant,required"`
	User   string `header:"X-Cs-User,required"`
}
