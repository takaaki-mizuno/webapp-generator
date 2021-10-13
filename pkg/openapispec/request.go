package openapispec

// Request ...
type Request struct {
	Path                      string
	PackageName               string
	RouteNameSpace            string
	PathFrameworkPresentation string
	Method                    string
	MethodCamel               string
	HandlerName               string
	HandlerFileName           string
	Description               string
	AddParamsForTest          string
	RequestSchemaName         Name
	Parameters                []*Parameter
	Responses                 []*Response
}
