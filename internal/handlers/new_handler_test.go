package handlers

import (
	"github.com/omiselabs/opn-generator/config"
	"github.com/omiselabs/opn-generator/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_NewHandler(t *testing.T) {
	t.Run("Run new command without api spec", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		gitService := services.MockGitService{}
		userAPIService := services.MockUserAPIService{}
		handler := NewNewHandler(configInstance, &gitService, &userAPIService)
		err := handler.Execute("test", ".", "")
		assert.Nil(t, err)
	})

	t.Run("Run new command with api spec", func(t *testing.T) {
		configInstance, _ := config.LoadConfig()
		gitService := services.MockGitService{}
		userAPIService := services.MockUserAPIService{}
		handler := NewNewHandler(configInstance, &gitService, &userAPIService)
		err := handler.Execute("test", ".", "test.yaml")
		assert.Nil(t, err)
	})
}
