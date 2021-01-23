package open_api_spec

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/inflection"
	"github.com/stoewer/go-strcase"
	"net/url"
	"strings"
)

func Parse(filePath string, namespace string, projectName string) (*API, error) {
	rootNamespace := "v1"
	data := API{
		FilePath:       filePath,
		BasePath:       "/",
		APINameSpace:   namespace,
		ProjectName:    projectName,
		Schemas:        map[string]*Schema{},
		RouteNameSpace: rootNamespace,
	}
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(filePath)
	if err != nil {
		return nil, err
	}
	if len(swagger.Servers) > 0 {
		elements, err := url.Parse(swagger.Servers[0].URL)
		if err == nil {
			data.BasePath = elements.Path
		}
	}
	parseComponents(swagger.Components, &data)

	for path, pathItem := range swagger.Paths {
		for method, operation := range pathItem.Operations() {
			request := Request{
				Path:           path,
				Method:         strings.ToUpper(method),
				MethodCamel:    strcase.UpperCamelCase(method),
				Description:    operation.Description,
				RouteNameSpace: rootNamespace,
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
					request.RequestSchemaName = generateName(requestSchema.Schema.Value.Title)
				}
			}
			for statusCode, schemaObject := range operation.Responses {
				responseSchema := schemaObject.Value.Content.Get("application/json")
				if responseSchema != nil {
					schema, ok := data.Schemas[responseSchema.Schema.Value.Title]
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

func parseComponents(components openapi3.Components, api *API) {
	for _, schemaRef := range components.Schemas {
		specSchema := schemaRef.Value
		if specSchema == nil {
			continue
		}
		if specSchema.Type != "object" {
			continue
		}
		schemaObject := Schema{
			Name:        specSchema.Title,
			Description: specSchema.Description,
		}
		requiredMap := map[string]bool{}
		for _, requiredColumn := range specSchema.Required {
			requiredMap[requiredColumn] = true
		}
		for name, property := range specSchema.Properties {
			_, required := requiredMap[name]
			switch property.Value.Type {
			case "array":
				item := property.Value.Items.Value
				schemaObject.Properties = append(schemaObject.Properties, &Property{
					Name:          name,
					Type:          property.Value.Type,
					Description:   property.Value.Description,
					ArrayItemType: item.Type,
					ArrayItemName: item.Title,
					Required:      required,
				})
			case "object":
				schemaObject.Properties = append(schemaObject.Properties, &Property{
					Name:        name,
					Type:        property.Value.Type,
					Description: property.Value.Description,
					Reference:   property.Value.Title,
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
		api.Schemas[specSchema.Title] = &schemaObject
	}
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
