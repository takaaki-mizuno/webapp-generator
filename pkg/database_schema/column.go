package database_schema

type Column struct {
	Name         string
	DataType     string
	Primary      bool
	DefaultValue string
	Nullable     bool
}
