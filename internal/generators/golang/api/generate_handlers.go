package api

import (
	"os"
	"strings"

	"github.com/takaaki-mizuno/webapp-generator/pkg/openapispec"
	"github.com/takaaki-mizuno/webapp-generator/pkg/template"
)

// GenerateHandlers ...
func GenerateHandlers(api *openapispec.API, path string) error {
	err := template.Replace(
		"api",
		"route",
		"route.tmpl",
		path,
		strings.Join([]string{"cmd", "app", "main.go"}, string(os.PathSeparator)),
		api,
	)
	if err != nil {
		return err
	}
	err = template.Generate(
		"api",
		"handler_struct.tmpl",
		path,
		strings.Join([]string{"internal", "http", api.APINameSpace, "handlers", "handler.go"}, string(os.PathSeparator)),
		api,
	)
	if err != nil {
		return err
	}
	for _, request := range api.Requests {
		err = template.Generate(
			"api",
			"handler.tmpl",
			path,
			strings.Join([]string{"internal", "http", api.APINameSpace, "handlers", request.HandlerFileName + "_handler.go"}, string(os.PathSeparator)),
			request,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"api",
			"handler_test.tmpl",
			path,
			strings.Join([]string{"internal", "http", api.APINameSpace, "handlers", request.HandlerFileName + "_handler_test.go"}, string(os.PathSeparator)),
			request,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
