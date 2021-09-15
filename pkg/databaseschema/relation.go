package databaseschema

// Relation ...
type Relation struct {
	Entity           *Entity
	Column           *Column
	RelationType     string
	ObjectName       string
	MultipleEntities bool
}
