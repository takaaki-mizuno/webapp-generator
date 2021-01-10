package database_schema

import (
	"github.com/jinzhu/inflection"
	"io/ioutil"
	"regexp"
	"strings"
)

func Parse(filePath string, projectName string) (*Schema, error) {
	data := Schema{
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
			Name:       name,
			HasDecimal: false,
			HasJson:    false,
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
				entityObject.Columns = append(entityObject.Columns, &Column{
					Name:     name,
					DataType: dataType,
					Primary:  primary,
					Nullable: nullable,
				})
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
		if relation[2] == "|" {
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
			Entity: rightTable,
			Column: nil,
		}
		multipleEntities := false
		if rightColumnIndex > -1 {
			leftRelation.Column = rightTable.Columns[rightColumnIndex]
			leftRelation.RelationType = "belongsTo"
		} else {
			leftRelation.Column = leftTable.Columns[leftColumnIndex]
			if rightRelationMany {
				leftRelation.RelationType = "hasMany"
				multipleEntities = true
			} else {
				leftRelation.RelationType = "hasOne"
			}
		}
		rightRelation := Relation{
			Entity:           leftTable,
			Column:           nil,
			MultipleEntities: multipleEntities,
		}
		if leftColumnIndex > -1 {
			rightRelation.Column = leftTable.Columns[leftColumnIndex]
			rightRelation.RelationType = "belongsTo"
		} else {
			rightRelation.Column = rightTable.Columns[rightColumnIndex]
			if leftRelationMany {
				rightRelation.RelationType = "hasMany"
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
