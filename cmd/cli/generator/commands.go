package generator

import (
	"github.com/omiselabs/opn-generator/cmd"
	"github.com/omiselabs/opn-generator/config"
	"github.com/omiselabs/opn-generator/internal/services"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// New ... make
func New(command *cobra.Command, args []string) error {
	container := cmd.BuildContainer()
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

	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = gitServiceInterface.DownloadBoilerplate(currentPath, args[0])
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}