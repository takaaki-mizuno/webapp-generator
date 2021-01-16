package open_api_spec

type Schema struct {
	Name        string
	Description string
	ObjectName  string
	Properties  []*Property
}
