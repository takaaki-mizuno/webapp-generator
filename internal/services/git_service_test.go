package services

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/takaaki-mizuno/webapp-generator/config"
)

func Test_gitService_DownloadBoilerplate(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		gitService := NewGitService(configInstance)

		projectName := "test"

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		err = gitService.DownloadBoilerplate(currentPath, projectName)
		assert.Nil(t, err)

		os.RemoveAll(currentPath + string(os.PathSeparator) + projectName)
	})

	t.Run("path error", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		gitService := NewGitService(configInstance)

		projectName := "test"

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		err = gitService.DownloadBoilerplate("", projectName)
		assert.NotNil(t, err)

		os.Remove(currentPath + string(os.PathSeparator) + "go-boilerplate-main.zip")
	})

	t.Run("missing project name", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		gitService := NewGitService(configInstance)

		projectName := ""

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		err = gitService.DownloadBoilerplate(currentPath, projectName)
		assert.NotNil(t, err)

		os.RemoveAll(currentPath + string(os.PathSeparator) + "go-boilerplate-main")
		os.Remove(currentPath + string(os.PathSeparator) + "go-boilerplate-main.zip")
	})
}
