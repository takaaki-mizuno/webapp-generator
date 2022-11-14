package databaseschema

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/stoewer/go-strcase"
)

// Parse ...
func Parse(filePath string, projectName string, organizationName string) (*Schema, error) {
	data := Schema{
		FilePath:           filePath,
		ProjectName:        projectName,
		OrganizationName:   organizationName,
		PrimaryKeyDataType: "int64",
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	cleanContent := removeComment(string(content))
	entityRegex := regexp.MustCompile(`(?m)entity "([^"]+)" {([^}]+)}`)
	relationRegex := regexp.MustCompile(`([a-z0-9_]+)\s+([}|])o--o([{|])\s+([a-z0-9_]+)`)
	columnRegex := regexp.MustCompile(`(\*|-)?\s*([a-z0-9_]+)\s*:\s*(\S+)`)
	entities := entityRegex.FindAllStringSubmatch(cleanContent, -1)
	relations := relationRegex.FindAllStringSubmatch(cleanContent, -1)

	for _, entity := range entities {
		entityName := entity[1]
		entityObject := Entity{
			Name:          generateName(entityName),
			HasDecimal:    false,
			HasJSON:       false,
			UseSoftDelete: false,
		}
		columns := strings.Split(strings.TrimSpace(entity[2]), "\n")
		for _, column := range columns {
			foundColumns := columnRegex.FindAllStringSubmatch(column, 1)
			if len(foundColumns) > 0 {
				primary := false
				nullable := true
				name := strings.ToLower(foundColumns[0][2])
				dataType := strings.ToLower(foundColumns[0][3])
				defaultValue := ""
				if name == "created_at" || name == "updated_at" {
					continue
				}
				if foundColumns[0][1] == "*" {
					nullable = false
				}
				if name == "id" {
					primary = true
					if dataType != "uuid" && dataType != "text" {
						dataType = "bigserial"
					} else {
						defaultValue = "uuid_generate_v4()"
						data.PrimaryKeyDataType = "string"
					}
				}
				columnObject := &Column{
					TableName:    entityObject.Name,
					Name:         generateName(name),
					DataType:     dataType,
					Primary:      primary,
					Nullable:     nullable,
					DefaultValue: defaultValue,
				}
				columnObject.APIReturnable = checkAPIReturnable(columnObject)
				columnObject.APIUpdatable = checkAPIUpdatable(columnObject)
				columnObject.APIType = getAPIType(columnObject)
				if name == "id" {
					columnObject.IsCommonColumn = true
				} else {
					columnObject.IsCommonColumn = false
				}
				entityObject.Columns = append(entityObject.Columns, columnObject)
				if name == "id" {
					entityObject.PrimaryKey = columnObject
				}
				if strings.HasPrefix(dataType, "decimal") || strings.HasPrefix(dataType, "numeric") {
					entityObject.HasDecimal = true
				}
				if strings.HasPrefix(dataType, "json") {
					entityObject.HasJSON = true
				}
				if name == "deleted_at" {
					columnObject.IsSystemUseColumn = true
					entityObject.UseSoftDelete = true
				} else {
					columnObject.IsSystemUseColumn = false
				}
			}
		}
		data.Entities = append(data.Entities, &entityObject)
	}

	for _, relation := range relations {
		leftTableName := relation[1]
		rightTableName := relation[4]
		leftRelationMany := true
		rightRelationMany := true
		if relation[2] == "|" {
			leftRelationMany = false
		}
		if relation[3] == "|" {
			rightRelationMany = false
		}
		leftTableIndex := findEntityIndex(leftTableName, &data)
		rightTableIndex := findEntityIndex(rightTableName, &data)
		if leftTableIndex == -1 || rightTableIndex == -1 {
			continue
		}
		leftTable := data.Entities[leftTableIndex]
		rightTable := data.Entities[rightTableIndex]
		rightColumnIndex := findRelationColumnIndex(leftTableName, rightTable)
		leftColumnIndex := findRelationColumnIndex(rightTableName, leftTable)
		if (leftColumnIndex == -1 && rightColumnIndex == -1) ||
			(leftColumnIndex != -1 && rightColumnIndex != -1) {
			continue
		}
		leftRelation := Relation{
			Entity:           rightTable,
			Column:           nil,
			MultipleEntities: false,
		}
		if leftColumnIndex > -1 {
			leftRelation.Column = leftTable.Columns[leftColumnIndex]
			leftRelation.RelationType = "belongsTo"
			leftTable.Columns[leftColumnIndex].RelationTableName = rightTable.Name
		} else {
			leftRelation.Column = rightTable.Columns[rightColumnIndex]
			if rightRelationMany {
				leftRelation.RelationType = "hasMany"
				leftRelation.MultipleEntities = true
			} else {
				leftRelation.RelationType = "hasOne"
			}
		}
		rightRelation := Relation{
			Entity:           leftTable,
			Column:           nil,
			MultipleEntities: false,
		}
		if rightColumnIndex > -1 {
			rightRelation.Column = rightTable.Columns[rightColumnIndex]
			rightRelation.RelationType = "belongsTo"
			rightTable.Columns[rightColumnIndex].RelationTableName = leftTable.Name
		} else {
			rightRelation.Column = leftTable.Columns[leftColumnIndex]
			if leftRelationMany {
				rightRelation.RelationType = "hasMany"
				leftRelation.MultipleEntities = true
			} else {
				rightRelation.RelationType = "hasOne"
			}
		}
		data.Entities[leftTableIndex].Relations = append(data.Entities[leftTableIndex].Relations, &leftRelation)
		data.Entities[rightTableIndex].Relations = append(data.Entities[rightTableIndex].Relations, &rightRelation)
	}
	return &data, nil
}

func findEntityIndex(name string, schema *Schema) int {
	for index, entity := range schema.Entities {
		if entity.Name.Original == name {
			return index
		}
	}
	return -1
}

func findRelationColumnIndex(name string, table *Entity) int {
	columnName := inflection.Singular(name) + "_id"
	for index, column := range table.Columns {
		if column.Name.Original == columnName {
			return index
		}
	}
	return -1
}

func removeComment(content string) string {
	commentRegex := regexp.MustCompile(`(?ms)\/'.+?'\/`)
	return commentRegex.ReplaceAllString(content, "")
}

func checkAPIReturnable(column *Column) bool {
	return true
}

func checkAPIUpdatable(column *Column) bool {
	if column.Name.Original == "id" || column.Name.Original == "created_at" || column.Name.Original == "updated_at" || column.Name.Original == "deleted_at" {
		return false
	}
	return true
}

func getAPIType(column *Column) string {
	if strings.HasPrefix(column.DataType, "decimal") || strings.HasPrefix(column.DataType, "numeric") {
		return "number"
	}
	switch column.DataType {
	case "int":
		return "integer"
	case "bigint":
		return "string"
	case "text":
		return "string"
	case "boolean":
		return "boolean"
	case "float":
		return "number"
	case "jsonb":
		return "string"
	case "timestamp":
		return "integer"
	}
	return "string"
}

func generateName(name string) Name {
	singular := inflection.Singular(name)
	plural := inflection.Plural(name)
	return Name{
		Original: name,
		Singular: NameForm{
			Camel: strcase.LowerCamelCase(singular),
			Title: strcase.UpperCamelCase(singular),
			Snake: strcase.SnakeCase(singular),
			Kebab: strcase.KebabCase(singular),
		},
		Plural: NameForm{
			Camel: strcase.LowerCamelCase(plural),
			Title: strcase.UpperCamelCase(plural),
			Snake: strcase.SnakeCase(plural),
			Kebab: strcase.KebabCase(plural),
		},
	}
}
