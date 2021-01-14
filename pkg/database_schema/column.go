package database_schema

type Column struct {
	Name          string
	DataType      string
	Primary       bool
	DefaultValue  string
	Nullable      bool
	ObjectName    string
	ObjectType    string
	APIReturnable bool
	APIUpdatable  bool
	APIType       string
}
