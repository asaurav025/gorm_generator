package gormgenerator

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Generate generate model files according to their name
func Generate(modelPath, fileName string) {
	modelList := []Model{}
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal([]byte(yamlFile), &modelList)
	modelTemp, err := ioutil.ReadFile("modelGenerator/model.go.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("model").Parse(string(modelTemp))
	if err != nil {
		panic(err)
	}

	for _, modelT := range modelList {
		modelFilePath := fmt.Sprintf("%s/%s.go", modelPath, modelT.Name)
		f, err := os.Create(modelFilePath)
		if err != nil {
			panic(err)
		}
		tmpl.Execute(f, &modelT)
		f.Close()
	}

}
