package services

import (
	"github.com/omiselabs/opn-generator/config"
)

// DatabaseServiceInterface ...
type DatabaseServiceInterface interface {
	GenerateDatabase(path string, databaseDefinitionPath string, language string, projectName string) error
	GenerateAdminAPI(path string, databaseDefinitionPath string, language string, projectName string) error
}

// DatabaseService ...
type DatabaseService struct {
	config *config.Config
}

func (service *DatabaseService) GenerateDatabase(path string, apiDefinitionPath string, language string, projectName string) error {
	return nil
}

func (service *DatabaseService) GenerateAdminAPI(path string, apiDefinitionPath string, language string, projectName string) error {
	return nil
}

// NewDatabaseService ...
func NewDatabaseService(
	config *config.Config,
) DatabaseServiceInterface {
	return &DatabaseService{
		config: config,
	}
}
