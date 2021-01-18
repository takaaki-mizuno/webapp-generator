package database

import (
	"github.com/jinzhu/inflection"
	"github.com/opn-ooo/opn-generator/pkg/database_schema"
	"github.com/opn-ooo/opn-generator/pkg/template"
	"os"
	"strings"
)

func GenerateModels(schema *database_schema.Schema, path string) error {
	for _, entity := range schema.Entities {
		err := template.Generate(
			"database",
			"model.tmpl",
			path,
			strings.Join([]string{"internal", "models", inflection.Singular(entity.Name) + ".go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"database",
			"model_test.tmpl",
			path,
			strings.Join([]string{"internal", "models", inflection.Singular(entity.Name) + "_test.go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
