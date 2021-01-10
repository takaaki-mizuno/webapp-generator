package database_schema

type Schema struct {
	ProjectName string
	PackageName string
	Entities    []*Entity
}
