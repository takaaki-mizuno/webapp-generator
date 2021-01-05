package services

import (
	"github.com/omiselabs/opn-generator/config"
	"github.com/omiselabs/opn-generator/pkg/open_api_spec"
)

// UserAPIServiceInterface ...
type UserAPIServiceInterface interface {
	GenerateUserAPI(path string, apiDefinitionPath string, language string) error
}

// UserAPIService ...
type UserAPIService struct {
	config *config.Config
}

func (service *UserAPIService) GenerateUserAPI(path string, apiDefinitionPath string, language string) error {
	_, err := open_api_spec.Parse(apiDefinitionPath)
	if err != nil {
		return err
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
