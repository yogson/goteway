package main

import (
	"goteway/server"
	"goteway/utils"
)

var configFile = "api.yaml"

func main() {
	config := utils.ReadFile(configFile)
	srv := server.New().WithYamlConfig(config)
	srv.Run()
}
