package database

import (
	"github.com/opn-ooo/opn-generator/pkg/database_schema"
	"github.com/opn-ooo/opn-generator/pkg/template"
	"os"
	"strings"
)

func GenerateRepositories(schema *database_schema.Schema, path string) error {
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

func AddRepositoryToDIContainer(schema *database_schema.Schema, path string) error {
	err := template.Replace(
		"database",
		"repository",
		"repository_di_container.tmpl",
		path,
		strings.Join([]string{"cmd", "container.go"}, string(os.PathSeparator)),
		schema,
	)
	return err
}
