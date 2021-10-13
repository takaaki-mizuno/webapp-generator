package database

import (
	"os"
	"strings"

	"github.com/opn-ooo/opn-generator/pkg/databaseschema"
	"github.com/opn-ooo/opn-generator/pkg/template"
)

// GenerateAdminAPISpec ...
func GenerateAdminAPISpec(schema *databaseschema.Schema, path string) error {
	err := template.Generate(
		"database",
		"admin_api_spec.tmpl",
		path,
		strings.Join([]string{"docs", "admin_api.yaml"}, string(os.PathSeparator)),
		schema,
	)
	return err
}
