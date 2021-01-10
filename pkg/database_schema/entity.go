package database_schema

type Entity struct {
	Name        string
	ObjectName  string
	Columns     []*Column
	Relations   []*Relation
	Description string
	HasDecimal  bool
	HasJson     bool
	PackageName string
}
