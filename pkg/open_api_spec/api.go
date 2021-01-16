package open_api_spec

type API struct {
	FilePath       string
	ProjectName    string
	APINameSpace   string
	PackageName    string
	BasePath       string
	RouteNameSpace string
	Requests       []*Request
	Schemas        map[string]*Schema
}
