package open_api_spec

type Request struct {
	Path                      string
	PackageName               string
	PathFrameworkPresentation string
	Method                    string
	MethodCamel               string
	HandlerName               string
	HandlerFileName           string
	Description               string
	AddParamsForTest          string
	ProcessRequest            string
	BuildResponse             string
	RequestSchemaName         string
	Parameters                []*Parameter
	Responses                 []*Response
}
