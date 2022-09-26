package services

import (
	"github.com/takaaki-mizuno/webapp-generator/config"
	"github.com/takaaki-mizuno/webapp-generator/internal/generators"
	"github.com/takaaki-mizuno/webapp-generator/pkg/databaseschema"
)

// DatabaseServiceInterface ...
type DatabaseServiceInterface interface {
	GenerateDatabase(path string, databaseDefinitionPath string, language string, projectName string, organizationName string) error
	GenerateAdminAPI(path string, databaseDefinitionPath string, language string, projectName string, organizationName string) error
}

// DatabaseService ...
type DatabaseService struct {
	config *config.Config
}

// GenerateDatabase ...
func (service *DatabaseService) GenerateDatabase(path string, databaseDefinitionPath string, language string, projectName string, organizationName string) error {
	schema, err := databaseschema.Parse(databaseDefinitionPath, projectName, organizationName)
	if err != nil {
		return err
	}
	generator := generators.NewGenerator(language)
	if generator != nil {
		err = generator.GenerateEntityInformation(schema, path)
		if err != nil {
			return err
		}
	}
	return nil
}

// GenerateAdminAPI ...
func (service *DatabaseService) GenerateAdminAPI(path string, apiDefinitionPath string, language string, projectName string, organizationName string) error {
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
