package database

import (
	"github.com/opn-ooo/opn-generator/pkg/database_schema"
	"github.com/opn-ooo/opn-generator/pkg/template"
	"os"
	"strings"
)

func GenerateAdminAPISpec(schema *database_schema.Schema, path string) error {
	err := template.Generate(
		"database",
		"admin_api_spec.tmpl",
		path,
		strings.Join([]string{"docs", "admin_api.yaml"}, string(os.PathSeparator)),
		schema,
	)
	return err
}
