package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/takaaki-mizuno/webapp-generator/config"
	"github.com/takaaki-mizuno/webapp-generator/internal/services"
)

func TestHandler_NewHandler(t *testing.T) {
	t.Run("Run new command without api spec", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		gitService := services.MockGitService{}
		userAPIService := services.MockUserAPIService{}
		databaseService := services.MockDatabaseService{}
		handler := NewNewHandler(configInstance, &gitService, &userAPIService, &databaseService)
		err := handler.Execute("test", ".", "", "", "test", "go")
		assert.Nil(t, err)
	})

	t.Run("Run new command with api spec", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		gitService := services.MockGitService{}
		userAPIService := services.MockUserAPIService{}
		databaseService := services.MockDatabaseService{}
		handler := NewNewHandler(configInstance, &gitService, &userAPIService, &databaseService)
		err := handler.Execute("test", ".", "test.yaml", "", "test", "go")
		assert.Nil(t, err)
	})
}
