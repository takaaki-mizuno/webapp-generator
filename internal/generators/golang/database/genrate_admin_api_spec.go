package database

import (
	"os"
	"strings"

	"github.com/takaaki-mizuno/webapp-generator/pkg/databaseschema"
	"github.com/takaaki-mizuno/webapp-generator/pkg/template"
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
