package open_api_spec

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
)

func Parse(filePath string) (*API, error) {
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(filePath)
	if err != nil {
		return nil, err
	}
	for path, pathItem := range swagger.Paths {
		fmt.Println(path)
		for method, _ := range pathItem.Operations() {
			fmt.Println(method)

		}
	}

	return nil, nil
}
