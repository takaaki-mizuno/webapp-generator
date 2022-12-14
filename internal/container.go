package internal

import (
	"github.com/takaaki-mizuno/webapp-generator/config"
	"github.com/takaaki-mizuno/webapp-generator/internal/handlers"
	"github.com/takaaki-mizuno/webapp-generator/internal/services"
	"go.uber.org/dig"
)

// BuildContainer ... setup DI container
func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.LoadConfig)
	_ = container.Provide(services.NewGitService)
	_ = container.Provide(services.NewUserAPIService)
	_ = container.Provide(services.NewDatabaseService)

	_ = container.Provide(handlers.NewNewHandler)

	return container
}
