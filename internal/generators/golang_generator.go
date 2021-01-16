package generators

import (
	"fmt"
	"github.com/jinzhu/inflection"
	"github.com/omiselabs/opn-generator/pkg/database_schema"
	"github.com/omiselabs/opn-generator/pkg/files"
	"github.com/omiselabs/opn-generator/pkg/open_api_spec"
	"github.com/omiselabs/opn-generator/pkg/template"
	"github.com/stoewer/go-strcase"
	"os"
	"strings"
	"time"
)

// GitServiceInterface ...
type GolangGenerator struct {
}

func (generator *GolangGenerator) GenerateRequestInformation(api *open_api_spec.API, path string) error {
	err := files.CopyFile(api.FilePath, strings.Join([]string{path, "docs", "user_api.yaml"}, string(os.PathSeparator)))
	if err != nil {
		return err
	}
	err = buildRequestLanguageSpecificInfo(api)
	if err != nil {
		return err
	}
	err = generateRequestRelatedFiles(api, path)
	if err != nil {
		return err
	}
	return err
}

func (generator *GolangGenerator) GenerateEntityInformation(schema *database_schema.Schema, path string) error {
	err := files.CopyFile(schema.FilePath, strings.Join([]string{path, "docs", "database.puml"}, string(os.PathSeparator)))
	if err != nil {
		return err
	}
	err = buildEntityLanguageSpecificInfo(schema)
	if err != nil {
		return err
	}
	err = generateModelRelatedFiles(schema, path)
	return err
}

func buildRequestLanguageSpecificInfo(api *open_api_spec.API) error {
	api.PackageName = "github.com/omiselabs/" + api.ProjectName

	api.RouteNameSpace, _ = buildRouteNameSpace(api)

	for schemaIndex, schema := range api.Schemas {
		api.Schemas[schemaIndex].ObjectName = strcase.UpperCamelCase(schema.Name)
		for propertyIndex, property := range schema.Properties {
			objectType := "String"
			api.Schemas[schemaIndex].Properties[propertyIndex].ObjectName = strcase.UpperCamelCase(property.Name)
			switch property.Type {
			case "string":
				objectType = "string"
			case "integer":
				objectType = "int64"
			case "number":
				objectType = "float64"
			case "boolean":
				objectType = "bool"
			case "array":
				arrayType := strcase.UpperCamelCase(property.ArrayItemName)
				switch property.ArrayItemType {
				case "string":
					arrayType = "string"
				case "integer":
					arrayType = "int64"
				case "number":
					arrayType = "float64"
				case "boolean":
					arrayType = "bool"
				}
				objectType = "[]" + arrayType
			case "object":
				objectType = strcase.UpperCamelCase(property.Reference)
			}
			api.Schemas[schemaIndex].Properties[propertyIndex].ObjectType = objectType
		}
	}

	for index, request := range api.Requests {
		api.Requests[index].HandlerName, _ = buildHandlerName(request)
		api.Requests[index].HandlerFileName = strcase.SnakeCase(api.Requests[index].HandlerName)
		api.Requests[index].PathFrameworkPresentation, _ = buildPathPresentation(request)
		api.Requests[index].PackageName = api.PackageName
	}

	return nil
}

func buildRouteNameSpace(api *open_api_spec.API) (string, error) {
	if api.BasePath == "/" {
		return api.APINameSpace, nil
	}
	elements := strings.Split(api.BasePath, "/")
	name := ""
	for _, element := range elements {
		if element != "" {
			name = name + strcase.UpperCamelCase(element)
		}
	}
	return strcase.LowerCamelCase(name), nil
}

func buildHandlerName(request *open_api_spec.Request) (string, error) {
	method := strcase.UpperCamelCase(strings.ToLower(request.Method))
	if request.Path == "/" {
		return "Index" + method, nil
	}
	elements := strings.Split(request.Path, "/")
	name := ""
	for index, element := range elements {
		if element == "" {
			continue
		}
		if !strings.HasPrefix(element, "{") {
			if index+1 < len(elements) {
				name = name + strcase.UpperCamelCase(inflection.Singular(element))
			} else {
				name = name + strcase.UpperCamelCase(element)
			}
		}
	}
	name = name + method
	return name, nil
}

func buildPathPresentation(request *open_api_spec.Request) (string, error) {
	elements := strings.Split(request.Path, "/")
	var result []string
	for _, element := range elements {
		if strings.HasPrefix(element, "{") {
			result = append(result, ":"+strings.TrimLeft(strings.TrimLeft(element, "}"), "{"))
		} else {
			result = append(result, element)
		}
	}

	return strings.Join(result, "/"), nil
}

func generateRequestRelatedFiles(api *open_api_spec.API, path string) error {
	err := generateHandlersAndTests(api, path)
	if err != nil {
		return err
	}
	err = generateRequestsAndResponses(api, path)
	return err
}

func generateHandlersAndTests(api *open_api_spec.API, path string) error {
	for _, request := range api.Requests {
		err := template.Generate(
			"api",
			"handler.tmpl",
			path,
			strings.Join([]string{"internal", "http", api.APINameSpace, "handlers", request.HandlerFileName + ".go"}, string(os.PathSeparator)),
			request,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"api",
			"handler_test.tmpl",
			path,
			strings.Join([]string{"internal", "http", api.APINameSpace, "handlers", request.HandlerFileName + "_test.go"}, string(os.PathSeparator)),
			request,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateRequestsAndResponses(api *open_api_spec.API, path string) error {
	for _, request := range api.Requests {
		for _, response := range request.Responses {
			responseSchemaName := response.Schema.Name
			responseObjectPath := strings.Join([]string{"internal", "http", api.APINameSpace, "responses", responseSchemaName + ".go"}, string(os.PathSeparator))
			schema, ok := api.Schemas[responseSchemaName]
			if files.Exists(responseObjectPath) == false && ok {
				err := template.Generate(
					"api",
					"response.tmpl",
					path,
					responseObjectPath,
					schema,
				)
				if err != nil {
					return err
				}
			}
		}
		if request.RequestSchemaName != "" {
			requestObjectPath := strings.Join([]string{"internal", "http", api.APINameSpace, "requests", request.RequestSchemaName + ".go"}, string(os.PathSeparator))
			schema, ok := api.Schemas[request.RequestSchemaName]
			if files.Exists(requestObjectPath) == false && ok {
				err := template.Generate(
					"api",
					"response.tmpl",
					path,
					requestObjectPath,
					schema,
				)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func buildEntityLanguageSpecificInfo(schema *database_schema.Schema) error {
	schema.PackageName = "github.com/omiselabs/" + schema.ProjectName
	for index, entity := range schema.Entities {
		schema.Entities[index].ObjectName = buildModelObjectName(entity)
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

func generateModelRelatedFiles(schema *database_schema.Schema, path string) error {
	err := generateModelsAndTests(schema, path)
	if err != nil {
		return err
	}
	err = generateRepositoriesAndTests(schema, path)
	if err != nil {
		return err
	}
	err = generateMigrations(schema, path)
	if err != nil {
		return err
	}
	err = generateAdminAPISpec(schema, path)
	return err
}

func generateModelsAndTests(schema *database_schema.Schema, path string) error {
	for _, entity := range schema.Entities {
		err := template.Generate(
			"database",
			"model.tmpl",
			path,
			strings.Join([]string{"internal", "models", inflection.Singular(entity.Name) + ".go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"database",
			"model_test.tmpl",
			path,
			strings.Join([]string{"internal", "models", inflection.Singular(entity.Name) + "_test.go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateRepositoriesAndTests(schema *database_schema.Schema, path string) error {
	for _, entity := range schema.Entities {
		err := template.Generate(
			"database",
			"repository.tmpl",
			path,
			strings.Join([]string{"internal", "repositories", inflection.Singular(entity.Name) + ".go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"database",
			"repository_test.tmpl",
			path,
			strings.Join([]string{"internal", "repositories", inflection.Singular(entity.Name) + "_test.go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"database",
			"repository_mock.tmpl",
			path,
			strings.Join([]string{"internal", "repositories", inflection.Singular(entity.Name) + "_mock.go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateMigrations(schema *database_schema.Schema, path string) error {
	currentTime := time.Now()
	prefix := currentTime.Format("20060102150405")

	for index, entity := range schema.Entities {
		filename := fmt.Sprintf("%s_%02d_create_%s", prefix, index, entity.Name)
		err := template.Generate(
			"database",
			"migration_up.tmpl",
			path,
			strings.Join([]string{"database", "migrations", filename + ".up.sql"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"database",
			"migration_down.tmpl",
			path,
			strings.Join([]string{"database", "migrations", filename + ".down.sql"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateAdminAPISpec(schema *database_schema.Schema, path string) error {
	err := template.Generate(
		"database",
		"admin_api_spec.tmpl",
		path,
		strings.Join([]string{"docs", "admin_api.yaml"}, string(os.PathSeparator)),
		schema,
	)
	return err
}
