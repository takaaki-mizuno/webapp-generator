package api

import (
	"github.com/opn-ooo/opn-generator/pkg/open_api_spec"
	"github.com/opn-ooo/opn-generator/pkg/template"
	"os"
	"strings"
)

func GenerateHandlers(api *open_api_spec.API, path string) error {
	for _, request := range api.Requests {
		err := template.Generate(
			"api",
			"handler.tmpl",
			path,
			strings.Join([]string{"internal", "http", api.APINameSpace, "handlers", request.HandlerFileName + ".go"}, string(os.PathSeparator)),
			request,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"api",
			"handler_test.tmpl",
			path,
			strings.Join([]string{"internal", "http", api.APINameSpace, "handlers", request.HandlerFileName + "_test.go"}, string(os.PathSeparator)),
			request,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
