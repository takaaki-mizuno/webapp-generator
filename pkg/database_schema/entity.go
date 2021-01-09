package database_schema

type Entity struct {
	Name      string
	Columns   []*Column
	Relations []*Relation
}
