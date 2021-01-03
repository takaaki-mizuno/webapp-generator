package handlers

import (
	"github.com/omiselabs/opn-generator/config"
	"github.com/omiselabs/opn-generator/internal"
	"github.com/omiselabs/opn-generator/internal/services"
	"github.com/omiselabs/opn-generator/pkg/open_api_spec"
	"log"
)

// New ... make new project
func NewHandler(projectName string, targetPath string, apiDefinitionPath string) error {
	container := internal.BuildContainer()
	var configInstance *config.Config
	var gitServiceInterface services.GitServiceInterface

	if err := container.Invoke(func(
		_configInstance *config.Config,
		_gitServiceInterface services.GitServiceInterface,
	) {
		gitServiceInterface = _gitServiceInterface
		configInstance = _configInstance
	}); err != nil {
		log.Fatal(err)
		return err
	}

	err := gitServiceInterface.DownloadBoilerplate(targetPath, projectName)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if apiDefinitionPath != "" {
		_, err := open_api_spec.Parse(apiDefinitionPath)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
