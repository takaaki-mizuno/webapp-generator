package template

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"

	htmlTemplate "text/template"
)

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

func Replace(templateType string, placeName string, templateFileName string, projectBasePath string, destinationFilePath string, data interface{}) error {
	templateFilePath := getTemplatePath(templateType, templateFileName, projectBasePath)

	templateInstance := htmlTemplate.Must(htmlTemplate.ParseFiles(templateFilePath))
	buffer := &bytes.Buffer{}

	err := templateInstance.Execute(buffer, data)
	if err != nil {
		return err
	}

	filePointer, err := os.Open(projectBasePath + string(os.PathSeparator) + destinationFilePath)
	if err != nil {
		return err
	}
	defer filePointer.Close()
	reader := bufio.NewReaderSize(filePointer, 4096)
	var lines []string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		lines = append(lines, string(line))
	}
	startString := "{{ REPLACE " + placeName + " }}"
	endString := "{{ REPLACE END " + placeName + " }}"
	var result []string
	replacing := false
	for _, line := range lines {
		if !replacing {
			if strings.Contains(line, startString) {
				result = append(result, line)
				replacing = true
			} else {
				result = append(result, line)
			}
		} else {
			if strings.Contains(line, endString) {
				result = append(result, buffer.String())
				result = append(result, line)
				replacing = false
			}
		}
	}

	err = ioutil.WriteFile(projectBasePath+string(os.PathSeparator)+destinationFilePath, []byte(strings.Join(result, "\n")), 0644)
	if err != nil {
		return err
	}

	return nil
}

func getTemplatePath(templateType string, fileName string, projectBasePath string) string {
	return projectBasePath + string(os.PathSeparator) + "templates" + string(os.PathSeparator) + templateType + string(os.PathSeparator) + fileName
}
