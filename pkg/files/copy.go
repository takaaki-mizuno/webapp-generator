package files

import (
	"io/ioutil"
	"os"
)

// CopyFile ...
func CopyFile(sourcePath string, destinationPath string) error {
	data, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(destinationPath, data, 0644)
	return err
}

// Exists ...
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
