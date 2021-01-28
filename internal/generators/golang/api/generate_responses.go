package api

import (
	"fmt"
	"github.com/opn-ooo/opn-generator/pkg/files"
	"github.com/opn-ooo/opn-generator/pkg/open_api_spec"
	"github.com/opn-ooo/opn-generator/pkg/template"
	"os"
	"strings"
)

func GenerateResponses(api *open_api_spec.API, path string) error {
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
	}
	return nil
}

func generateResponseStruct(name string, api *open_api_spec.API, path string) error {
	schema, ok := api.Schemas[name]
	if !ok {
		return fmt.Errorf("response %s not found", name)
	}
	responseObjectPath := strings.Join([]string{"internal", "http", api.APINameSpace, "responses", name + ".go"}, string(os.PathSeparator))
	if files.Exists(responseObjectPath) {
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
