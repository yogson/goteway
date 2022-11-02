package main

import (
	"module-rest/server"
	"module-rest/utils"
)

var configFile = "api.yaml"

func main() {
	config := utils.ReadFile(configFile)
	srv := server.New().WithYamlConfig(config)
	srv.Run()
}
