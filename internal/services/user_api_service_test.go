package services

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opn-ooo/opn-generator/config"
)

func Test_userAPIService_GenerateUserAPI(t *testing.T) {
	t.Run("API spec path error", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		userAPIService := NewUserAPIService(configInstance)

		// projectName := "test"

		currentPath, err := os.Getwd()
		assert.Nil(t, err)

		apiDefinitionPath, err := filepath.Abs("../../test/user_api.yaml")
		assert.Nil(t, err)
		assert.NotNil(t, apiDefinitionPath)

		language := "golang"
		projectName := "test"

		err = userAPIService.GenerateUserAPI(currentPath+string(os.PathSeparator)+projectName, apiDefinitionPath, language, projectName)
		assert.NotNil(t, err)
	})
}
