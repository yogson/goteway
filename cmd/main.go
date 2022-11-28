package main

import (
	"goteway/server"
	"goteway/utils"
	_ "goteway/handlers"
	_ "goteway/middleware"
)

var version string

var configFile = utils.GetEnv("GW-CONFIG", "../api.yaml")

func main() {
	config := utils.ReadFile(configFile)
	srv := server.New().WithYamlConfig(config)
	srv.Run()
}
