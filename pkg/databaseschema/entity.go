package databaseschema

// Entity ...
type Entity struct {
	Name        Name
	Columns     []*Column
	Relations   []*Relation
	PrimaryKey  *Column
	Description string
	HasDecimal  bool
	HasJson     bool
	PackageName string
}
