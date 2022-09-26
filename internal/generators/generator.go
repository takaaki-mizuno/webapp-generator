package generators

import (
	"github.com/takaaki-mizuno/webapp-generator/pkg/databaseschema"
	"github.com/takaaki-mizuno/webapp-generator/pkg/openapispec"
)

// GeneratorInterface ...
type GeneratorInterface interface {
	GenerateRequestInformation(api *openapispec.API, path string) error
	GenerateEntityInformation(schema *databaseschema.Schema, path string) error
}

// NewGenerator ...
func NewGenerator(language string) GeneratorInterface {
	if language == "golang" {
		return &GolangGenerator{}
	}

	return nil
}
