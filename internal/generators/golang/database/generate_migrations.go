package database

import (
	"fmt"
	"github.com/opn-ooo/opn-generator/pkg/database_schema"
	"github.com/opn-ooo/opn-generator/pkg/template"
	"os"
	"strings"
	"time"
)

func GenerateMigrations(schema *database_schema.Schema, path string) error {
	currentTime := time.Now()
	prefix := currentTime.Format("200601021504")

	for index, entity := range schema.Entities {
		filename := fmt.Sprintf("%s%02d_create_%s", prefix, index, entity.Name)
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
