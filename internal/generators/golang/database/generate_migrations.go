package database

import (
	"fmt"
	"os"
	"strings"

	"github.com/opn-ooo/opn-generator/pkg/databaseschema"
	"github.com/opn-ooo/opn-generator/pkg/template"
)

// GenerateMigrations ...
func GenerateMigrations(schema *databaseschema.Schema, path string) error {

	for index, entity := range schema.Entities {
		filename := fmt.Sprintf("%06d_create_%s", index+2, entity.Name.Plural.Snake)
		err := template.Generate(
			"database",
			"migration_up.tmpl",
			path,
			strings.Join([]string{"database", "migrations", filename + ".up.sql"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"database",
			"migration_down.tmpl",
			path,
			strings.Join([]string{"database", "migrations", filename + ".down.sql"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
