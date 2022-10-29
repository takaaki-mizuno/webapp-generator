package services

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/takaaki-mizuno/webapp-generator/config"
)

func Test_userAPIService_GenerateUserAPI(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		userAPIService := NewUserAPIService(configInstance)

		projectName := "test"
		language := "golang"
		organizationName := "test"

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		tempProjectPath := currentPath + string(os.PathSeparator) + projectName

		apiDefinitionPath, err := filepath.Abs("../../test/user_api.yaml")
		assert.Nil(t, err)
		assert.NotNil(t, apiDefinitionPath)

		gitService := NewGitService(configInstance)
		err = gitService.DownloadBoilerplate(currentPath, projectName, "go")
		assert.Nil(t, err)

		err = userAPIService.GenerateUserAPI(tempProjectPath, apiDefinitionPath, language, projectName, organizationName)
		assert.Nil(t, err)

		os.RemoveAll(tempProjectPath)
	})

	t.Run("empty project path makes failure", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		userAPIService := NewUserAPIService(configInstance)

		projectName := "test"
		organizationName := "test"

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		apiDefinitionPath, err := filepath.Abs("../../test/user_api.yaml")
		assert.Nil(t, err)
		assert.NotNil(t, apiDefinitionPath)

		gitService := NewGitService(configInstance)
		err = gitService.DownloadBoilerplate(currentPath, projectName, "go")
		assert.Nil(t, err)

		language := "golang"
		tempProjectPath := currentPath + string(os.PathSeparator) + projectName

		err = userAPIService.GenerateUserAPI("", apiDefinitionPath, language, projectName, organizationName)
		assert.NotNil(t, err)

		os.RemoveAll(tempProjectPath)
	})

	t.Run("empty api spec path makes failure", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		userAPIService := NewUserAPIService(configInstance)

		projectName := "test"
		organizationName := "test"

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		gitService := NewGitService(configInstance)
		err = gitService.DownloadBoilerplate(currentPath, projectName, "go")
		assert.Nil(t, err)

		language := "golang"
		tempProjectPath := currentPath + string(os.PathSeparator) + projectName

		err = userAPIService.GenerateUserAPI(tempProjectPath, "", language, projectName, organizationName)
		assert.NotNil(t, err)

		os.RemoveAll(tempProjectPath)
	})

	t.Run("empty project name makes failure", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		userAPIService := NewUserAPIService(configInstance)

		projectName := ""
		language := "golang"
		organizationName := "test"

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		tempProjectPath := currentPath + string(os.PathSeparator) + projectName

		apiDefinitionPath, err := filepath.Abs("../../test/user_api.yaml")
		assert.Nil(t, err)
		assert.NotNil(t, apiDefinitionPath)

		gitService := NewGitService(configInstance)
		err = gitService.DownloadBoilerplate(currentPath, "test", "go")
		assert.Nil(t, err)

		err = userAPIService.GenerateUserAPI(tempProjectPath, apiDefinitionPath, language, projectName, organizationName)
		assert.NotNil(t, err)

		os.RemoveAll(tempProjectPath + string(os.PathSeparator) + "test")
	})
}
