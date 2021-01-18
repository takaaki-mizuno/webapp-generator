package database_schema

type Entity struct {
	Name             string
	SingularName     string
	ObjectName       string
	ObjectPluralName string
	Columns          []*Column
	Relations        []*Relation
	Description      string
	HasDecimal       bool
	HasJson          bool
	PackageName      string
}
