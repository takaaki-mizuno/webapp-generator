package files

import "io/ioutil"

func CopyFile(sourcePath string, destinationPath string) error {
	data, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(destinationPath, data, 0644)
	return err
}
