package database

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/opn-ooo/opn-generator/pkg/databaseschema"
	"github.com/opn-ooo/opn-generator/pkg/template"
)

// GenerateMigrations ...
func GenerateMigrations(schema *databaseschema.Schema, path string) error {
	currentTime := time.Now()
	prefix := currentTime.Format("200601021504")

	for index, entity := range schema.Entities {
		filename := fmt.Sprintf("%s%02d_create_%s", prefix, index, entity.Name.Plural.Snake)
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
