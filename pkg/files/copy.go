package files

import (
	"io/ioutil"
	"os"
)

func CopyFile(sourcePath string, destinationPath string) error {
	data, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(destinationPath, data, 0644)
	return err
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
