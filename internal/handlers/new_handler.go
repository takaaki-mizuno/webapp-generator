package handlers

import (
	"github.com/omiselabs/opn-generator/config"
	"github.com/omiselabs/opn-generator/internal/services"
	"log"
)

// NewHandler ... handler for new command
type NewHandler struct {
	config         *config.Config
	gitService     services.GitServiceInterface
	userAPIService services.UserAPIServiceInterface
}

// Execute ... make new project
func (handler *NewHandler) Execute(projectName string, targetPath string, apiDefinitionPath string) error {
	err := handler.gitService.DownloadBoilerplate(targetPath, projectName)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if apiDefinitionPath != "" {
		err = handler.userAPIService.GenerateUserAPI(targetPath, apiDefinitionPath, "golang")
	}

	return nil
}

// NewNewHandler ...
func NewNewHandler(
	config *config.Config,
	gitService services.GitServiceInterface,
	userAPIService services.UserAPIServiceInterface,
) *NewHandler {
	return &NewHandler{
		config:         config,
		gitService:     gitService,
		userAPIService: userAPIService,
	}
}
