package services

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cavaliercoder/grab"

	"github.com/takaaki-mizuno/webapp-generator/config"
)

// GitServiceInterface ...
type GitServiceInterface interface {
	DownloadBoilerplate(path string, projectName string) error
}

// GitService ...
type GitService struct {
	config *config.Config
}

// DownloadBoilerplate ...
func (service *GitService) DownloadBoilerplate(path string, projectName string) error {
	zipFilePath, err := downloadFile(service.config.Boilerplate.URL, path)
	if err != nil {
		return err
	}
	err = unzip(zipFilePath, path)
	if err != nil {
		return err
	}
	_, zipFileName := filepath.Split(zipFilePath)
	slice := strings.Split(zipFileName, ".")
	err = os.Rename(slice[0], projectName)
	if err != nil {
		return err
	}

	err = os.Remove(zipFilePath)
	if err != nil {
		return err
	}
	return nil
}

func downloadFile(targetURL string, directoryPath string) (string, error) {
	response, err := grab.Get(directoryPath, targetURL)
	if err != nil && response.HTTPResponse.StatusCode != 200 {
		log.Fatal(err)
		return "", err
	}
	return response.Filename, nil
}

// Source: https://github.com/artdarek/go-unzip/blob/master/unzip.go
func unzip(zipFilePath string, destinationFilePath string) error {

	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	err = os.MkdirAll(destinationFilePath, 0755)
	if err != nil {
		return err
	}

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(destinationFilePath, f.Name)
		if !strings.HasPrefix(path, filepath.Clean(destinationFilePath)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: Illegal file path", path)
		}

		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(path, f.Mode())
		} else {
			_ = os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

// NewGitService ...
func NewGitService(
	config *config.Config,
) GitServiceInterface {
	return &GitService{
		config: config,
	}
}
