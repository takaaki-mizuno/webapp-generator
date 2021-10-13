package database

import (
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/stoewer/go-strcase"

	"github.com/opn-ooo/opn-generator/pkg/databaseschema"
)

// BuildLanguageSpecificInfo ...
func BuildLanguageSpecificInfo(schema *databaseschema.Schema) error {
	schema.PackageName = "github.com/opn-ooo/" + schema.ProjectName
	for index := range schema.Entities {
		schema.Entities[index].PackageName = schema.PackageName
		for columnIndex, column := range schema.Entities[index].Columns {
			schema.Entities[index].Columns[columnIndex].ObjectName = buildColumnObjectName(column)
			schema.Entities[index].Columns[columnIndex].ObjectType = buildColumnObjectType(column)
			schema.Entities[index].Columns[columnIndex].APIObjectType = buildColumnAPIObjectType(column)
		}
		for relationIndex, relation := range schema.Entities[index].Relations {
			schema.Entities[index].Relations[relationIndex].ObjectName = buildRelationObjectName(relation)
		}
	}
	return nil
}

func buildColumnObjectName(column *databaseschema.Column) string {
	name := strcase.UpperCamelCase(column.Name.Original)
	if strings.HasSuffix(name, "Id") {
		name = name[:len(name)-1] + "D"
	}
	return name
}

func buildColumnObjectType(column *databaseschema.Column) string {
	dataType := strings.ToLower(column.DataType)
	if strings.HasPrefix(dataType, "decimal") {
		return "decimal.Decimal"
	}
	switch dataType {
	case "text":
		return "string"
	case "int":
		return "int32"
	case "bigserial":
		return "int64"
	case "bigint":
		return "int64"
	case "timestamp":
		return "time.Time"
	case "boolean":
		return "bool"
	case "jsonb":
		return "datatypes.JSON"
	}

	return "string"
}

func buildColumnAPIObjectType(column *databaseschema.Column) string {
	dataType := strings.ToLower(column.DataType)
	if strings.HasPrefix(dataType, "decimal") {
		return "string"
	}
	switch dataType {
	case "text":
		return "string"
	case "int":
		return "int32"
	case "bigserial":
		return "int64"
	case "bigint":
		return "int64"
	case "timestamp":
		return "int64"
	case "boolean":
		return "bool"
	case "jsonb":
		return "string"
	}

	return "string"
}

func buildRelationObjectName(relation *databaseschema.Relation) string {
	if relation.MultipleEntities {
		return strcase.UpperCamelCase(relation.Entity.Name.Original)
	}
	return strcase.UpperCamelCase(inflection.Singular(relation.Entity.Name.Original))
}
