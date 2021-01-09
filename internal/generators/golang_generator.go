package generators

import (
	"github.com/jinzhu/inflection"
	"github.com/omiselabs/opn-generator/pkg/open_api_spec"
	"github.com/omiselabs/opn-generator/pkg/template"
	"github.com/stoewer/go-strcase"
	"log"
	"os"
	"strings"
)

// GitServiceInterface ...
type GolangGenerator struct {
}

func (generator *GolangGenerator) GenerateRequestInformation(api *open_api_spec.API, path string) error {
	err := buildRequestLanguageSpecificInfo(api)
	if err != nil {
		return err
	}
	err = generateRequestRelatedFiles(api, path)
	return err
}

func buildRequestLanguageSpecificInfo(api *open_api_spec.API) error {
	api.PackageName = "github.com/omiselabs/" + api.ProjectName

	api.RouteNameSpace, _ = buildRouteNameSpace(api)
	for index, request := range api.Requests {
		api.Requests[index].HandlerName, _ = buildHandlerName(request)
		api.Requests[index].HandlerFileName = strcase.SnakeCase(api.Requests[index].HandlerName)
		api.Requests[index].PathFrameworkPresentation, _ = buildPathPresentation(request)
		api.Requests[index].PackageName = api.PackageName
	}

	return nil
}

func buildRouteNameSpace(api *open_api_spec.API) (string, error) {
	if api.BasePath == "/" {
		return api.APINameSpace, nil
	}
	elements := strings.Split(api.BasePath, "/")
	name := ""
	for _, element := range elements {
		if element != "" {
			name = name + strcase.UpperCamelCase(element)
		}
	}
	return strcase.LowerCamelCase(name), nil
}

func buildHandlerName(request *open_api_spec.Request) (string, error) {
	method := strcase.UpperCamelCase(strings.ToLower(request.Method))
	if request.Path == "/" {
		return "Index" + method, nil
	}
	elements := strings.Split(request.Path, "/")
	name := ""
	for index, element := range elements {
		if element == "" {
			continue
		}
		if !strings.HasPrefix(element, "{") {
			if index+1 < len(elements) {
				name = name + strcase.UpperCamelCase(inflection.Singular(element))
			} else {
				name = name + strcase.UpperCamelCase(element)
			}
		}
	}
	name = name + method
	return name, nil
}

func buildPathPresentation(request *open_api_spec.Request) (string, error) {
	elements := strings.Split(request.Path, "/")
	var result []string
	for _, element := range elements {
		if strings.HasPrefix(element, "{") {
			result = append(result, ":"+strings.TrimLeft(strings.TrimLeft(element, "}"), "{"))
		} else {
			result = append(result, element)
		}
	}

	return strings.Join(result, "/"), nil
}

func generateRequestRelatedFiles(api *open_api_spec.API, path string) error {
	err := generateHandlerAndTests(api, path)
	if err != nil {
		return err
	}
	return nil
}

func generateHandlerAndTests(api *open_api_spec.API, path string) error {
	for _, request := range api.Requests {
		err := template.Generate(
			"api",
			"handler.tmpl",
			path,
			strings.Join([]string{"internal", "http", api.APINameSpace, "handlers", request.HandlerFileName + ".go"}, string(os.PathSeparator)),
			request,
		)
		if err != nil {
			log.Fatal(err)
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
