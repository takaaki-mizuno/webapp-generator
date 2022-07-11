package services

import (
	"github.com/opn-ooo/opn-generator/config"
	"github.com/opn-ooo/opn-generator/internal/generators"
	"github.com/opn-ooo/opn-generator/pkg/openapispec"
)

// UserAPIServiceInterface ...
type UserAPIServiceInterface interface {
	GenerateUserAPI(path string, apiDefinitionPath string, language string, projectName string, organizationName string) error
}

// UserAPIService ...
type UserAPIService struct {
	config *config.Config
}

// GenerateUserAPI ...
func (service *UserAPIService) GenerateUserAPI(path string, apiDefinitionPath string, language string, projectName string, organizationName string) error {
	api, err := openapispec.Parse(apiDefinitionPath, "app", projectName, organizationName)
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
