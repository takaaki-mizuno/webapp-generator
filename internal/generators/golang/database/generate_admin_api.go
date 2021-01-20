package database

import (
	"github.com/opn-ooo/opn-generator/pkg/database_schema"
	"github.com/opn-ooo/opn-generator/pkg/template"
	"os"
	"strings"
)

func GenerateAdminAPIHandlers(schema *database_schema.Schema, path string) error {
	err := template.Generate(
		"database",
		"admin_api_handler.tmpl",
		path,
		strings.Join([]string{"internal", "http", "admin", "handlers", "handler.go"}, string(os.PathSeparator)),
		schema,
	)
	if err != nil {
		return err
	}
	for _, entity := range schema.Entities {
		err = template.Generate(
			"database",
			"admin_api_entity_handler.tmpl",
			path,
			strings.Join([]string{"internal", "http", "admin", "handlers", entity.Name.Singular.Snake + "_handler.go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateAdminAPIRequests(schema *database_schema.Schema, path string) error {
	for _, entity := range schema.Entities {
		err := template.Generate(
			"database",
			"admin_api_update_request.tmpl",
			path,
			strings.Join([]string{"internal", "http", "admin", "requests", entity.Name.Singular.Snake + "_update.go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateAdminAPIResponse(schema *database_schema.Schema, path string) error {
	for _, entity := range schema.Entities {
		err := template.Generate(
			"database",
			"admin_api_list_response.tmpl",
			path,
			strings.Join([]string{"internal", "http", "admin", "responses", entity.Name.Plural.Snake + ".go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
		err = template.Generate(
			"database",
			"admin_api_single_response.tmpl",
			path,
			strings.Join([]string{"internal", "http", "admin", "responses", entity.Name.Singular.Snake + ".go"}, string(os.PathSeparator)),
			entity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
