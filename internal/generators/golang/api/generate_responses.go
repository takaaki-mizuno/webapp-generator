package api

import (
	"fmt"
	"os"
	"strings"

	"github.com/stoewer/go-strcase"

	"github.com/opn-ooo/opn-generator/pkg/files"
	"github.com/opn-ooo/opn-generator/pkg/openapispec"
	"github.com/opn-ooo/opn-generator/pkg/template"
)

// GenerateResponses ...
func GenerateResponses(api *openapispec.API, path string) error {
	for _, request := range api.Requests {
		for _, response := range request.Responses {
			err := generateResponseStruct(response.Schema.Name, api, path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func generateResponseStruct(name string, api *openapispec.API, path string) error {
	schema, ok := api.Schemas[name]
	if !ok {
		return fmt.Errorf("response %s not found", name)
	}
	fileName := strcase.SnakeCase(name)
	responseObjectPath := strings.Join([]string{"internal", "http", api.APINameSpace, "responses", fileName + ".go"}, string(os.PathSeparator))
	if files.Exists(path + string(os.PathSeparator) + responseObjectPath) {
		return nil
	}

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
	for _, property := range schema.Properties {
		if property.Reference != "" {
			err := generateResponseStruct(property.Reference, api, path)
			if err != nil {
				return err
			}
		} else if property.ArrayItemName != "" {
			err := generateResponseStruct(property.ArrayItemName, api, path)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
