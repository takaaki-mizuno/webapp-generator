package api

import (
	"github.com/opn-ooo/opn-generator/pkg/files"
	"github.com/opn-ooo/opn-generator/pkg/open_api_spec"
	"github.com/opn-ooo/opn-generator/pkg/template"
	"os"
	"strings"
)

func GenerateRequests(api *open_api_spec.API, path string) error {
	for _, request := range api.Requests {
		if request.RequestSchemaName.Original != "" {
			requestObjectPath := strings.Join([]string{"internal", "http", api.APINameSpace, "requests", request.RequestSchemaName.Original + ".go"}, string(os.PathSeparator))
			schema, ok := api.Schemas[request.RequestSchemaName.Original]
			if files.Exists(requestObjectPath) == false && ok {
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
			}
		}
	}
	return nil
}
