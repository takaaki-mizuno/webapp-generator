package services

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opn-ooo/opn-generator/config"
)

func Test_databaseService_GenerateDatabase(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		databaseService := NewDatabaseService(configInstance)

		projectName := "test"
		language := "golang"

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		tempProjectPath := currentPath + string(os.PathSeparator) + projectName

		apiDefinitionPath, err := filepath.Abs("../../test/database.puml")
		assert.Nil(t, err)
		assert.NotNil(t, apiDefinitionPath)

		gitService := NewGitService(configInstance)
		err = gitService.DownloadBoilerplate(currentPath, projectName)
		assert.Nil(t, err)

		err = databaseService.GenerateDatabase(tempProjectPath, apiDefinitionPath, language, projectName)
		assert.Nil(t, err)

		os.RemoveAll(tempProjectPath)
	})

	t.Run("empty project path makes failure", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		databaseService := NewDatabaseService(configInstance)

		projectName := "test"
		language := "golang"

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		tempProjectPath := currentPath + string(os.PathSeparator) + projectName

		apiDefinitionPath, err := filepath.Abs("../../test/database.puml")
		assert.Nil(t, err)
		assert.NotNil(t, apiDefinitionPath)

		gitService := NewGitService(configInstance)
		err = gitService.DownloadBoilerplate(currentPath, projectName)
		assert.Nil(t, err)

		err = databaseService.GenerateDatabase("", apiDefinitionPath, language, projectName)
		assert.NotNil(t, err)

		os.RemoveAll(tempProjectPath)
	})

	t.Run("empty database schema path makes failure", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		databaseService := NewDatabaseService(configInstance)

		projectName := "test"
		language := "golang"

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		tempProjectPath := currentPath + string(os.PathSeparator) + projectName

		gitService := NewGitService(configInstance)
		err = gitService.DownloadBoilerplate(currentPath, projectName)
		assert.Nil(t, err)

		err = databaseService.GenerateDatabase(tempProjectPath, "", language, projectName)
		assert.NotNil(t, err)

		os.RemoveAll(tempProjectPath)
	})

	t.Run("empty project name makes failure", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		databaseService := NewDatabaseService(configInstance)

		projectName := "test"
		language := "golang"

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		tempProjectPath := currentPath + string(os.PathSeparator) + projectName

		apiDefinitionPath, err := filepath.Abs("../../test/database.puml")
		assert.Nil(t, err)
		assert.NotNil(t, apiDefinitionPath)

		gitService := NewGitService(configInstance)
		err = gitService.DownloadBoilerplate(currentPath, projectName)
		assert.Nil(t, err)

		err = databaseService.GenerateDatabase(currentPath, apiDefinitionPath, language, "")
		assert.NotNil(t, err)

		os.RemoveAll(tempProjectPath)
	})
}
