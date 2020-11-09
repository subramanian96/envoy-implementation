package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

// Exporting the yaml configuration.
type ConfigYaml struct {
	Network struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
}

func main() {
    fmt.Println("Parsing the YAML file")

    var fileName string
    flag.StringVar(&fileName, "f", "", "YAML file to parse.")
    flag.Parse()

    if fileName == "" {
        fmt.Println("Specify the yaml file by using -f option")
        return
    }

    yamlFile, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Printf("Error fetching YAML file: %s\n", err)
        return
    }

    var configYaml ConfigYaml
    err = yaml.Unmarshal(yamlFile, &configYaml)
    if err != nil {
        fmt.Printf("Error parsing YAML file: %s\n", err)
    }

    fmt.Printf("Config: %v\n", configYaml)
}
