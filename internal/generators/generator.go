package generators

import (
	"github.com/opn-ooo/opn-generator/pkg/databaseschema"
	"github.com/opn-ooo/opn-generator/pkg/openapispec"
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
