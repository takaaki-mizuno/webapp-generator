package handlers

import (
	"github.com/omiselabs/opn-generator/config"
	"github.com/omiselabs/opn-generator/internal/services"
	"log"
	"os"
)

// NewHandler ... handler for new command
type NewHandler struct {
	config          *config.Config
	gitService      services.GitServiceInterface
	userAPIService  services.UserAPIServiceInterface
	databaseService services.DatabaseServiceInterface
}

// Execute ... make new project
func (handler *NewHandler) Execute(projectName string, targetPath string, apiDefinitionPath string, databaseDefinitionPath string) error {
	err := handler.gitService.DownloadBoilerplate(targetPath, projectName)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if apiDefinitionPath != "" {
		err = handler.userAPIService.GenerateUserAPI(targetPath+string(os.PathSeparator)+projectName, apiDefinitionPath, "golang", projectName)
		if err != nil {
			return err
		}
	}

	if databaseDefinitionPath != "" {
		err = handler.databaseService.GenerateDatabase(targetPath+string(os.PathSeparator)+projectName, databaseDefinitionPath, "golang", projectName)
		if err != nil {
			return err
		}
		err = handler.databaseService.GenerateAdminAPI(targetPath+string(os.PathSeparator)+projectName, databaseDefinitionPath, "golang", projectName)
		if err != nil {
			return err
		}
	}

	return nil
}

// NewNewHandler ...
func NewNewHandler(
	config *config.Config,
	gitService services.GitServiceInterface,
	userAPIService services.UserAPIServiceInterface,
	databaseService services.DatabaseServiceInterface,
) *NewHandler {
	return &NewHandler{
		config:          config,
		gitService:      gitService,
		userAPIService:  userAPIService,
		databaseService: databaseService,
	}
}
