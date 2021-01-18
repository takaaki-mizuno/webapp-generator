package database

import (
	"github.com/jinzhu/inflection"
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
			strings.Join([]string{"internal", "repositories", inflection.Singular(entity.Name) + ".go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"database",
			"repository_test.tmpl",
			path,
			strings.Join([]string{"internal", "repositories", inflection.Singular(entity.Name) + "_test.go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"database",
			"repository_mock.tmpl",
			path,
			strings.Join([]string{"internal", "repositories", inflection.Singular(entity.Name) + "_mock.go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
