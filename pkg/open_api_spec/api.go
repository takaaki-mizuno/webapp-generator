package open_api_spec

type API struct {
	ProjectName    string
	APINameSpace   string
	PackageName    string
	BasePath       string
	RouteNameSpace string
	Requests       []*Request
}
