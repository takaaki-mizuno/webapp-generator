package database

import (
	"os"
	"strings"

	"github.com/takaaki-mizuno/webapp-generator/pkg/databaseschema"
	"github.com/takaaki-mizuno/webapp-generator/pkg/template"
)

// GenerateRepositories ...
func GenerateRepositories(schema *databaseschema.Schema, path string) error {
	err := template.Generate(
		"database",
		"base_repository.tmpl",
		path,
		strings.Join([]string{"internal", "repositories", "base_repository.go"}, string(os.PathSeparator)),
		schema,
	)
	if err != nil {
		return err
	}
	for _, entity := range schema.Entities {
		err := template.Generate(
			"database",
			"repository.tmpl",
			path,
			strings.Join([]string{"internal", "repositories", entity.Name.Singular.Snake + "_repository.go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"database",
			"repository_test.tmpl",
			path,
			strings.Join([]string{"internal", "repositories", entity.Name.Singular.Snake + "_repository_test.go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"database",
			"repository_mock.tmpl",
			path,
			strings.Join([]string{"internal", "repositories", entity.Name.Singular.Snake + "_repository_mock.go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddRepositoryToDIContainer ...
func AddRepositoryToDIContainer(schema *databaseschema.Schema, path string) error {
	err := template.Replace(
		"database",
		"repository",
		"repository_di_container.tmpl",
		path,
		strings.Join([]string{"cmd", "container.go"}, string(os.PathSeparator)),
		schema,
	)
	if err != nil {
		return err
	}
	err = template.ReplaceWithString(
		"repository_import",
		"\t\""+schema.PackageName+"/internal/repositories\"",
		path,
		strings.Join([]string{"cmd", "container.go"}, string(os.PathSeparator)),
	)
	return err
}
