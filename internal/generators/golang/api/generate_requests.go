package api

import (
	"fmt"
	"os"
	"strings"

	"github.com/stoewer/go-strcase"

	"github.com/takaaki-mizuno/webapp-generator/pkg/files"
	"github.com/takaaki-mizuno/webapp-generator/pkg/openapispec"
	"github.com/takaaki-mizuno/webapp-generator/pkg/template"
)

// GenerateRequests ...
func GenerateRequests(api *openapispec.API, path string) error {
	for _, request := range api.Requests {
		if request.RequestSchemaName.Original != "" {
			err := generateRequestStruct(request.RequestSchemaName.Original, api, path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func generateRequestStruct(name string, api *openapispec.API, path string) error {
	schema, ok := api.Schemas[name]
	if !ok {
		return fmt.Errorf("request %s not found", name)
	}
	fileName := strcase.SnakeCase(name)
	requestObjectPath := strings.Join([]string{"internal", "http", api.APINameSpace, "requests", fileName + ".go"}, string(os.PathSeparator))

	if files.Exists(path + string(os.PathSeparator) + requestObjectPath) {
		return nil
	}

	err := template.Generate(
		"api",
		"request.tmpl",
		path,
		requestObjectPath,
		schema,
	)
	if err != nil {
		return err
	}
	for _, property := range schema.Properties {
		if property.Reference != "" {
			err := generateRequestStruct(property.Reference, api, path)
			if err != nil {
				return err
			}
		} else if property.ArrayItemName != "" {
			err := generateRequestStruct(property.ArrayItemName, api, path)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
