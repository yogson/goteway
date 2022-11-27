package main

import (
	"goteway/server"
	"goteway/utils"
)

var version string

var configFile = utils.GetEnv("GW-CONFIG", "api-gateway.yaml")

func main() {
	config := utils.ReadFile(configFile)
	srv := server.New().WithYamlConfig(config)
	srv.Run()
}
