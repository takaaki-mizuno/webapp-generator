package open_api_spec

type Property struct {
	Name          string
	ObjectName    string
	Type          string
	ObjectType    string
	ArrayItemType string
	ArrayItemName string
	Description   string
	Reference     string
	Required      bool
}
