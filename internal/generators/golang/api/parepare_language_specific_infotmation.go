package api

import (
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/stoewer/go-strcase"

	"github.com/takaaki-mizuno/webapp-generator/pkg/openapispec"
)

// BuildLanguageSpecificInfo ...
func BuildLanguageSpecificInfo(api *openapispec.API) error {
	api.PackageName = "github.com/" + api.OrganizationName + "/" + api.ProjectName

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

func buildHandlerName(request *openapispec.Request) (string, error) {
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

func buildPathPresentation(request *openapispec.Request) (string, error) {
	elements := strings.Split(request.Path, "/")
	var result []string
	for _, element := range elements {
		if strings.HasPrefix(element, "{") {
			result = append(result, ":"+strings.TrimLeft(strings.TrimRight(element, "}"), "{"))
		} else {
			result = append(result, element)
		}
	}

	return strings.Join(result, "/"), nil
}
