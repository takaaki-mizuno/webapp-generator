package databaseschema

// Schema ...
type Schema struct {
	FilePath           string
	ProjectName        string
	PackageName        string
	PrimaryKeyDataType string
	Entities           []*Entity
}
