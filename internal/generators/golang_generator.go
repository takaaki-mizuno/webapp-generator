package generators

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	apiGenerator "github.com/takaaki-mizuno/webapp-generator/internal/generators/golang/api"
	databaseGenerator "github.com/takaaki-mizuno/webapp-generator/internal/generators/golang/database"
	"github.com/takaaki-mizuno/webapp-generator/pkg/databaseschema"
	"github.com/takaaki-mizuno/webapp-generator/pkg/files"
	"github.com/takaaki-mizuno/webapp-generator/pkg/openapispec"
)

// GolangGenerator ...
type GolangGenerator struct {
}

// GenerateRequestInformation ...
func (generator *GolangGenerator) GenerateRequestInformation(api *openapispec.API, path string) error {
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

// GenerateEntityInformation ...
func (generator *GolangGenerator) GenerateEntityInformation(schema *databaseschema.Schema, path string) error {
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

func generateAPIRelatedFiles(api *openapispec.API, path string) error {
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

func generateModelRelatedFiles(schema *databaseschema.Schema, path string) error {
	err := databaseGenerator.GenerateModels(schema, path)
	if err != nil {
		return err
	}
	err = databaseGenerator.GenerateRepositories(schema, path)
	if err != nil {
		return err
	}
	err = databaseGenerator.AddRepositoryToDIContainer(schema, path)
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
	err = databaseGenerator.GenerateAdminAPIRoutes(schema, path)
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
			code = strings.Replace(code, "github.com/takaaki-mizuno/go-boilerplate", packageName, -1)
			err = ioutil.WriteFile(path, []byte(code), 0644)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func formatCode(path string) error {
	fmt.Println(path)
	err := exec.Command("go", "fmt", path+"/...").Run()
	return err
}
