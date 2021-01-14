package open_api_spec

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stoewer/go-strcase"
	"net/url"
	"strings"
)

func Parse(filePath string, namespace string, projectName string) (*API, error) {
	data := API{
		FilePath:     filePath,
		BasePath:     "/",
		APINameSpace: namespace,
		ProjectName:  projectName,
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
	for path, pathItem := range swagger.Paths {
		for method, operation := range pathItem.Operations() {
			request := Request{
				Path:        path,
				Method:      strings.ToUpper(method),
				MethodCamel: strcase.UpperCamelCase(method),
				Description: operation.Description,
			}
			/*
				for statusCode, responseObject := range operation.Responses {
					for _, responseContent := range responseObject.Value.Content {
						parameter = Parameter{
							Name: responseContent.Schema.Value.Title,
							Type: responseContent.Schema.Value.Type,

						}



					}
					response := Response{
						StatusCode: statusCode,
					}
					if strings.HasPrefix(statusCode, "2") {
						// Success Case

					} else {
						// Error Case

					}
				}
			*/
			data.Requests = append(data.Requests, &request)
		}
	}

	return &data, nil
}
