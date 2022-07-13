package database

import (
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/stoewer/go-strcase"

	"github.com/opn-ooo/opn-generator/pkg/databaseschema"
)

// BuildLanguageSpecificInfo ...
func BuildLanguageSpecificInfo(schema *databaseschema.Schema) error {
	schema.PackageName = "github.com/" + schema.OrganizationName + "/" + schema.ProjectName
	for index := range schema.Entities {
		schema.Entities[index].PackageName = schema.PackageName
		for columnIndex, column := range schema.Entities[index].Columns {
			schema.Entities[index].Columns[columnIndex].ObjectName = buildColumnObjectName(column)
			schema.Entities[index].Columns[columnIndex].ObjectType = buildColumnObjectType(column)
			schema.Entities[index].Columns[columnIndex].APIObjectType = buildColumnAPIObjectType(column)
			schema.Entities[index].Columns[columnIndex].FakerType = buildFakerObjectType(column)
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
	case "float":
		return "float32"
	case "real":
		return "float32"
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
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "real":
		return "float32"
	case "json":
		return "string"
	case "jsonb":
		return "string"
	}

	return "string"
}

func buildFakerObjectType(column *databaseschema.Column) string {
	dataType := strings.ToLower(column.DataType)
	name := column.Name.Original

	switch dataType {
	case "text":
		if name == "id" || strings.HasSuffix(name, "_id") {
			return "uuid_digit"
		}
		if name == "url" || strings.HasSuffix(name, "_url") {
			return "url"
		}
		if name == "username" || name == "user_name" {
			return "username"
		}
		if name == "name" || strings.HasSuffix(name, "_name") {
			return "name"
		}
		if strings.HasSuffix(name, "email") {
			return "email"
		}
		if strings.HasSuffix(name, "description") {
			return "sentence"
		}
		return "word"
	case "real":
		return "amount"
	case "int":
		if strings.HasSuffix(name, "_at") {
			return "unix_time"
		}
		return "oneof: 500, 1000"
	case "bigserial":
		return "oneof: 500, 1000"
	case "bigint":
		if strings.HasSuffix(name, "_at") {
			return "unix_time"
		}
		return "oneof: 500, 1000"
	case "timestamp":
		return "unix_time"
	case "boolean":
		return "word"
	case "jsonb":
		return "word"
	}

	return "word"
}

func buildRelationObjectName(relation *databaseschema.Relation) string {
	if relation.MultipleEntities {
		return strcase.UpperCamelCase(relation.Entity.Name.Original)
	}
	return strcase.UpperCamelCase(inflection.Singular(relation.Entity.Name.Original))
}
