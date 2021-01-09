package generators

import "github.com/omiselabs/opn-generator/pkg/open_api_spec"

// GitServiceInterface ...
type GeneratorInterface interface {
	GenerateRequestInformation(api *open_api_spec.API, path string) error
}

func NewGenerator(language string) GeneratorInterface {
	if language == "golang" {
		return &GolangGenerator{}
	}

	return nil
}
