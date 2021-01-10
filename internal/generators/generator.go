package generators

import (
	"github.com/omiselabs/opn-generator/pkg/database_schema"
	"github.com/omiselabs/opn-generator/pkg/open_api_spec"
)

// GitServiceInterface ...
type GeneratorInterface interface {
	GenerateRequestInformation(api *open_api_spec.API, path string) error
	GenerateEntityInformation(schema *database_schema.Schema, path string) error
}

func NewGenerator(language string) GeneratorInterface {
	if language == "golang" {
		return &GolangGenerator{}
	}

	return nil
}
