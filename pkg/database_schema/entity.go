package database_schema

type Entity struct {
	Name        Name
	Columns     []*Column
	Relations   []*Relation
	Description string
	HasDecimal  bool
	HasJson     bool
	PackageName string
}
