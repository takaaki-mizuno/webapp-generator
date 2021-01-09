package database_schema

type Relation struct {
	Entity       *Entity
	Column       *Column
	RelationType string
}
