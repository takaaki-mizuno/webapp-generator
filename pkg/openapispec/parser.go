package openapispec

import (
	"net/url"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/inflection"
	"github.com/stoewer/go-strcase"
)

// Parse ...
func Parse(filePath string, namespace string, projectName string) (*API, error) {
	defaultRouteNamespace := namespace
	data := API{
		FilePath:       filePath,
		BasePath:       "/",
		APINameSpace:   namespace,
		ProjectName:    projectName,
		Schemas:        map[string]*Schema{},
		RouteNameSpace: defaultRouteNamespace,
	}
	swagger, err := openapi3.NewLoader().LoadFromFile(filePath)
	if err != nil {
		return nil, err
	}
	if len(swagger.Servers) > 0 {
		elements, err := url.Parse(swagger.Servers[0].URL)
		if err != nil {
			return nil, err
		}
		data.BasePath = elements.Path
	}
	parseComponents(swagger.Components, &data)

	data.RouteNameSpace = buildRouteNameSpace(data.BasePath, defaultRouteNamespace)

	for path, pathItem := range swagger.Paths {
		for method, operation := range pathItem.Operations() {
			request := Request{
				Path:           path,
				Method:         strings.ToUpper(method),
				MethodCamel:    strcase.UpperCamelCase(method),
				Description:    operation.Description,
				RouteNameSpace: data.RouteNameSpace,
			}
			// Parameters
			for _, parameterReference := range operation.Parameters {
				parameter := parameterReference.Value
				request.Parameters = append(request.Parameters, &Parameter{
					Name:     generateName(parameter.Name),
					In:       parameter.In,
					Required: parameter.Required,
				})
			}
			for _, parameterReference := range pathItem.Parameters {
				parameter := parameterReference.Value
				request.Parameters = append(request.Parameters, &Parameter{
					Name:     generateName(parameter.Name),
					In:       parameter.In,
					Required: parameter.Required,
				})
			}
			if operation.RequestBody != nil {
				requestSchema := operation.RequestBody.Value.Content.Get("application/json")
				if requestSchema != nil {
					if requestSchema.Schema.Ref != "" {
						requestName := getSchemaNameFromSchema(requestSchema.Schema.Ref, requestSchema.Schema.Value)
						request.RequestSchemaName = generateName(requestName)
					} else {
						requestSchemaName := strcase.SnakeCase(path) + "_" + strings.ToLower(method) + "_request"
						data.Schemas[requestSchemaName] = generateSchemaObject(requestSchemaName, requestSchema.Schema.Value)
						request.RequestSchemaName = generateName(requestSchemaName)
					}
				}
			}
			for statusCode, schemaObject := range operation.Responses {
				responseSchema := schemaObject.Value.Content.Get("application/json")

				if responseSchema != nil {
					responseName := getSchemaNameFromSchema(responseSchema.Schema.Ref, responseSchema.Schema.Value)
					schema, ok := data.Schemas[responseName]
					if ok {
						success := false
						if strings.HasPrefix(statusCode, "2") {
							success = true
						}
						request.Responses = append(request.Responses, &Response{
							StatusCode: statusCode,
							Schema:     schema,
							Success:    success,
						})
					}
				}
			}
			data.Requests = append(data.Requests, &request)
		}
	}

	return &data, nil
}

func buildRouteNameSpace(basePath string, defaultRouteNamespace string) string {
	if basePath == "/" || basePath == "" {
		return defaultRouteNamespace
	}

	elements := strings.Split(basePath, "/")
	name := ""
	for _, element := range elements {
		if element != "" {
			name = name + strcase.UpperCamelCase(element)
		}
	}
	return strcase.LowerCamelCase(name)
}

func parseComponents(components openapi3.Components, api *API) {
	for name, schemaRef := range components.Schemas {
		specSchema := schemaRef.Value
		if specSchema == nil {
			continue
		}
		if specSchema.Type != "object" {
			continue
		}
		schemaName := getSchemaNameFromSchema(name, schemaRef.Value)
		api.Schemas[schemaName] = generateSchemaObject(schemaName, schemaRef.Value)
	}
}

func generateSchemaObject(name string, schema *openapi3.Schema) *Schema {
	schemaObject := Schema{
		Name:        name,
		Description: schema.Description,
	}
	requiredMap := map[string]bool{}
	for _, requiredColumn := range schema.Required {
		requiredMap[requiredColumn] = true
	}
	for name, property := range schema.Properties {
		_, required := requiredMap[name]
		switch property.Value.Type {
		case "array":
			itemName := getSchemaNameFromSchema(property.Value.Items.Ref, property.Value.Items.Value)
			item := property.Value.Items.Value
			schemaObject.Properties = append(schemaObject.Properties, &Property{
				Name:          name,
				Type:          property.Value.Type,
				Description:   property.Value.Description,
				ArrayItemType: item.Type,
				ArrayItemName: itemName,
				Required:      required,
			})
		case "object":
			propertyName := getSchemaNameFromSchema(property.Ref, property.Value)
			schemaObject.Properties = append(schemaObject.Properties, &Property{
				Name:        name,
				Type:        property.Value.Type,
				Description: property.Value.Description,
				Reference:   propertyName,
				Required:    required,
			})
		default:
			schemaObject.Properties = append(schemaObject.Properties, &Property{
				Name:        name,
				Type:        property.Value.Type,
				Description: property.Value.Description,
				Required:    required,
			})
		}
	}
	return &schemaObject
}

func generateName(name string) Name {
	singular := inflection.Singular(name)
	plural := inflection.Plural(name)
	return Name{
		Original: name,
		Default: NameForm{
			Camel: strcase.LowerCamelCase(name),
			Title: strcase.UpperCamelCase(name),
			Snake: strcase.SnakeCase(name),
			Kebab: strcase.KebabCase(name),
		},
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

func getSchemaNameFromSchema(name string, schema *openapi3.Schema) string {
	if schema.Title != "" {
		return schema.Title
	}
	elements := strings.Split(name, "/")
	schemaName := strcase.SnakeCase(elements[len(elements)-1])
	return schemaName
}
