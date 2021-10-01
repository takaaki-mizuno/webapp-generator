package database

import (
	"fmt"
	"github.com/opn-ooo/opn-generator/pkg/database_schema"
	"github.com/opn-ooo/opn-generator/pkg/template"
	"os"
	"strings"
)

func GenerateMigrations(schema *database_schema.Schema, path string) error {

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
