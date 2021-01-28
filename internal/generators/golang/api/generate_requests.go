package api

import (
	"fmt"
	"github.com/opn-ooo/opn-generator/pkg/files"
	"github.com/opn-ooo/opn-generator/pkg/open_api_spec"
	"github.com/opn-ooo/opn-generator/pkg/template"
	"os"
	"strings"
)

func GenerateRequests(api *open_api_spec.API, path string) error {
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

func generateRequestStruct(name string, api *open_api_spec.API, path string) error {
	schema, ok := api.Schemas[name]
	if !ok {
		return fmt.Errorf("request %s not found", name)
	}
	requestObjectPath := strings.Join([]string{"internal", "http", api.APINameSpace, "requests", name + ".go"}, string(os.PathSeparator))

	if files.Exists(requestObjectPath) {
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
