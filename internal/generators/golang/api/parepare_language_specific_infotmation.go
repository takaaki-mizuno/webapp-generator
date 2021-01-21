package api

import (
	"github.com/jinzhu/inflection"
	"github.com/opn-ooo/opn-generator/pkg/open_api_spec"
	"github.com/stoewer/go-strcase"
	"strings"
)

func BuildLanguageSpecificInfo(api *open_api_spec.API) error {
	api.PackageName = "github.com/opn-ooo/" + api.ProjectName

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

func buildHandlerParam(api *open_api_spec.API, request *open_api_spec.Request) string {
data:
	for _, parameter := range request.Parameters {
		parameter.Name
	}
	if request.RequestSchemaName != "" {

	}

	return ""
}
