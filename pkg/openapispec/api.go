package openapispec

// API ...
type API struct {
	FilePath         string
	ProjectName      string
	APINameSpace     string
	PackageName      string
	OrganizationName string
	BasePath         string
	RouteNameSpace   string
	Requests         []*Request
	Schemas          map[string]*Schema
}
