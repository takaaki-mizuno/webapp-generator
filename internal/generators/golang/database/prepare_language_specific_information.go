package database

import (
	"github.com/jinzhu/inflection"
	"github.com/opn-ooo/opn-generator/pkg/database_schema"
	"github.com/stoewer/go-strcase"
	"strings"
)

func BuildLanguageSpecificInfo(schema *database_schema.Schema) error {
	schema.PackageName = "github.com/opn-ooo/" + schema.ProjectName
	for index, entity := range schema.Entities {
		schema.Entities[index].ObjectName = buildModelObjectName(entity)
		schema.Entities[index].ObjectPluralName = buildPluralModelObjectName(entity)
		schema.Entities[index].PackageName = schema.PackageName
		for columnIndex, column := range schema.Entities[index].Columns {
			schema.Entities[index].Columns[columnIndex].ObjectName = buildColumnObjectName(column)
			schema.Entities[index].Columns[columnIndex].ObjectType = buildColumnObjectType(column)
		}
		for relationIndex, relation := range schema.Entities[index].Relations {
			schema.Entities[index].Relations[relationIndex].ObjectName = buildRelationObjectName(relation)
		}
	}
	return nil
}

func buildModelObjectName(entity *database_schema.Entity) string {
	return strcase.UpperCamelCase(inflection.Singular(entity.Name))
}

func buildPluralModelObjectName(entity *database_schema.Entity) string {
	return strcase.UpperCamelCase(inflection.Plural(entity.Name))
}

func buildColumnObjectName(column *database_schema.Column) string {
	name := strcase.UpperCamelCase(column.Name)
	if strings.HasSuffix(name, "Id") {
		name = name[:len(name)-1] + "D"
	}
	return name
}

func buildColumnObjectType(column *database_schema.Column) string {
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
		return "postgres.Jsonb"
	}

	return "string"
}

func buildRelationObjectName(relation *database_schema.Relation) string {
	if relation.MultipleEntities {
		return strcase.UpperCamelCase(relation.Entity.Name)
	} else {
		return strcase.UpperCamelCase(inflection.Singular(relation.Entity.Name))
	}
}
