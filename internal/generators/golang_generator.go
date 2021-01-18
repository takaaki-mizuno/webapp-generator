package generators

import (
	apiGenerator "github.com/opn-ooo/opn-generator/internal/generators/golang/api"
	databaseGenerator "github.com/opn-ooo/opn-generator/internal/generators/golang/database"
	"io/ioutil"
	"path/filepath"

	"github.com/opn-ooo/opn-generator/pkg/database_schema"
	"github.com/opn-ooo/opn-generator/pkg/files"
	"github.com/opn-ooo/opn-generator/pkg/open_api_spec"
	"os"
	"strings"
)

// GitServiceInterface ...
type GolangGenerator struct {
}

func (generator *GolangGenerator) GenerateRequestInformation(api *open_api_spec.API, path string) error {
	err := files.CopyFile(api.FilePath, strings.Join([]string{path, "docs", "user_api.yaml"}, string(os.PathSeparator)))
	if err != nil {
		return err
	}
	err = apiGenerator.BuildLanguageSpecificInfo(api)
	if err != nil {
		return err
	}
	err = generateAPIRelatedFiles(api, path)
	if err != nil {
		return err
	}
	err = replacePackageName(path, api.PackageName)
	return err
}

func (generator *GolangGenerator) GenerateEntityInformation(schema *database_schema.Schema, path string) error {
	err := files.CopyFile(schema.FilePath, strings.Join([]string{path, "docs", "database.puml"}, string(os.PathSeparator)))
	if err != nil {
		return err
	}
	err = databaseGenerator.BuildLanguageSpecificInfo(schema)
	if err != nil {
		return err
	}
	err = generateModelRelatedFiles(schema, path)
	return err
}

func generateAPIRelatedFiles(api *open_api_spec.API, path string) error {
	err := apiGenerator.GenerateHandlers(api, path)
	if err != nil {
		return err
	}
	err = apiGenerator.GenerateRequests(api, path)
	if err != nil {
		return err
	}
	err = apiGenerator.GenerateResponses(api, path)
	return err
}

func generateModelRelatedFiles(schema *database_schema.Schema, path string) error {
	err := databaseGenerator.GenerateModels(schema, path)
	if err != nil {
		return err
	}
	err = databaseGenerator.GenerateRepositories(schema, path)
	if err != nil {
		return err
	}
	err = databaseGenerator.GenerateMigrations(schema, path)
	if err != nil {
		return err
	}
	err = databaseGenerator.GenerateAdminAPISpec(schema, path)
	if err != nil {
		return err
	}
	err = databaseGenerator.GenerateAdminAPIHandlers(schema, path)
	if err != nil {
		return err
	}
	err = databaseGenerator.GenerateAdminAPIRequests(schema, path)
	if err != nil {
		return err
	}
	err = databaseGenerator.GenerateAdminAPIResponse(schema, path)
	return err
}

func replacePackageName(path string, packageName string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "go.mod") {
			input, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			code := string(input)
			code = strings.Replace(code, "github.com/opn-ooo/go-boilerplate", packageName, -1)
			err = ioutil.WriteFile(path, []byte(code), 0644)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
