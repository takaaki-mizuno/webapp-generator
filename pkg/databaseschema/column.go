package databaseschema

// Column ...
type Column struct {
	Name          Name
	ObjectName    string
	DataType      string
	Primary       bool
	DefaultValue  string
	Nullable      bool
	ObjectType    string
	APIReturnable bool
	APIUpdatable  bool
	APIType       string
	APIObjectType string
}
