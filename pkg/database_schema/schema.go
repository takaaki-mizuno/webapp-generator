package database_schema

type Schema struct {
	FilePath           string
	ProjectName        string
	PackageName        string
	PrimaryKeyDataType string
	Entities           []*Entity
}
