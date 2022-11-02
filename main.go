package main

import (
	"rest-playground/server"
	"rest-playground/utils"
)

var configFile = "api.yaml"

func main() {
	config := utils.ReadFile(configFile)
	srv := server.New().WithYamlConfig(config)
	srv.Run()
}
