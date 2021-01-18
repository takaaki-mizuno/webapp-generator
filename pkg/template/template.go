package template

import (
	"bytes"
	"io/ioutil"
	"os"
)
import htmlTemplate "text/template"

func Generate(templateType string, templateFileName string, projectBasePath string, destinationFilePath string, data interface{}) error {
	templateFilePath := getTemplatePath(templateType, templateFileName, projectBasePath)

	templateInstance := htmlTemplate.Must(htmlTemplate.ParseFiles(templateFilePath))
	buffer := &bytes.Buffer{}

	err := templateInstance.Execute(buffer, data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(projectBasePath+string(os.PathSeparator)+destinationFilePath, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func getTemplatePath(templateType string, fileName string, projectBasePath string) string {
	return projectBasePath + string(os.PathSeparator) + "templates" + string(os.PathSeparator) + templateType + string(os.PathSeparator) + fileName
}
