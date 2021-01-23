package api

import (
	"github.com/opn-ooo/opn-generator/pkg/open_api_spec"
	"github.com/opn-ooo/opn-generator/pkg/template"
	"os"
	"strings"
)

func GenerateHandlers(api *open_api_spec.API, path string) error {
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
