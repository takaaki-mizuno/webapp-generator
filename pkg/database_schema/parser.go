package database_schema

import (
	"github.com/jinzhu/inflection"
	"io/ioutil"
	"regexp"
	"strings"
)

func Parse(filePath string, projectName string) (*Schema, error) {
	data := Schema{
		FilePath:    filePath,
		ProjectName: projectName,
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
		name := entity[1]
		entityObject := Entity{
			Name:         name,
			SingularName: inflection.Singular(name),
			HasDecimal:   false,
			HasJson:      false,
		}
		columns := strings.Split(strings.TrimSpace(entity[2]), "\n")
		for _, column := range columns {
			foundColumns := columnRegex.FindAllStringSubmatch(column, 1)
			if len(foundColumns) > 0 {
				primary := false
				nullable := false
				name = strings.ToLower(foundColumns[0][2])
				dataType := strings.ToLower(foundColumns[0][3])
				if name == "created_at" || name == "updated_at" {
					continue
				}
				if foundColumns[0][1] == "*" {
					nullable = true
				}
				if name == "id" {
					primary = true
					dataType = "bigserial"
				}
				columnObject := &Column{
					Name:     name,
					DataType: dataType,
					Primary:  primary,
					Nullable: nullable,
				}
				columnObject.APIReturnable = checkAPIReturnable(columnObject)
				columnObject.APIUpdatable = checkAPIUpdatable(columnObject)
				columnObject.APIType = getAPIType(columnObject)

				entityObject.Columns = append(entityObject.Columns, columnObject)
				if strings.HasPrefix(dataType, "decimal") || strings.HasPrefix(dataType, "numeric") {
					entityObject.HasDecimal = true
				}
				if strings.HasPrefix(dataType, "json") {
					entityObject.HasJson = true
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
		if entity.Name == name {
			return index
		}
	}
	return -1
}

func findRelationColumnIndex(name string, table *Entity) int {
	columnName := inflection.Singular(name) + "_id"
	for index, column := range table.Columns {
		if column.Name == columnName {
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
	if column.Name == "id" || column.Name == "created_at" || column.Name == "updated_at" {
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
	case "jsonb":
		return "string"
	case "timestamp":
		return "integer"
	}
	return "string"
}