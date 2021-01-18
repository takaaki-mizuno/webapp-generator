package services

import (
	"github.com/opn-ooo/opn-generator/config"
	"github.com/opn-ooo/opn-generator/internal/generators"
	"github.com/opn-ooo/opn-generator/pkg/open_api_spec"
)

// UserAPIServiceInterface ...
type UserAPIServiceInterface interface {
	GenerateUserAPI(path string, apiDefinitionPath string, language string, projectName string) error
}

// UserAPIService ...
type UserAPIService struct {
	config *config.Config
}

func (service *UserAPIService) GenerateUserAPI(path string, apiDefinitionPath string, language string, projectName string) error {
	api, err := open_api_spec.Parse(apiDefinitionPath, "app", projectName)
	if err != nil {
		return err
	}
	generator := generators.NewGenerator(language)
	if generator != nil {
		err = generator.GenerateRequestInformation(api, path)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewUserAPIService ...
func NewUserAPIService(
	config *config.Config,
) UserAPIServiceInterface {
	return &UserAPIService{
		config: config,
	}
}
