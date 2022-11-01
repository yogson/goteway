package endpoints

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Endpoints struct {
	Endpoints []Endpoint `yaml:"Endpoints"`
}

type MethodHandler struct {
	Method  string `yaml:"method"`
	Handler string `yaml:"handler"`
}

type Endpoint struct {
	Path    string          `yaml:"path"`
	Methods []MethodHandler `yaml:"methods"`
}

func (e *Endpoints) LoadFromYaml(path string) {
	yFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yFile, &e)
	if err != nil {
		log.Fatal(err)
	}
}
