# Gorm generator  
  
[![Build Status](https://travis-ci.org/asaurav025/gorm_generator.svg?branch=master)](https://travis-ci.org/asaurav025/gorm_generator)
[![Go Report Card](https://goreportcard.com/badge/github.com/asaurav025/gorm_generator)](https://goreportcard.com/report/github.com/asaurav025/gorm_generator)  
This golang generate gorm model based on yaml file input.  

## Steps to generate model  

- Write yaml file according to format specified in `sample.yaml`.  
- Multiple models can be defined in a single yaml file.  
- Import the package [`github.com/asaurav025/gorm_generator`](https://github.com/asaurav025/gorm_generator " Gorm generator")  
- Call `Generate()` function from your function, this function has two parameters `modelpath`: where to be model is to be generated, `filepath`: path/to/yamlfile  

## TODO  

- Add support to more filetype such as `json` and `xml`.  
- Look for better way to pass template to generate function, currently using utf-8 encoded format of template.  
