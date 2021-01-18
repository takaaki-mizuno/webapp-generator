package services

import (
	"github.com/opn-ooo/opn-generator/config"
	"github.com/opn-ooo/opn-generator/internal/generators"
	"github.com/opn-ooo/opn-generator/pkg/database_schema"
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

func (service *DatabaseService) GenerateDatabase(path string, databaseDefinitionPath string, language string, projectName string) error {
	schema, err := database_schema.Parse(databaseDefinitionPath, projectName)
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
