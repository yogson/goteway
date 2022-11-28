package main

import (
	"goteway/server"
	"goteway/utils"
	_ "goteway/handlers"
)

var version string

var configFile = utils.GetEnv("GW-CONFIG", "/Users/egor/GolandProjects/goteway/api.yaml")

func main() {
	config := utils.ReadFile(configFile)
	srv := server.New().WithYamlConfig(config)
	srv.Run()
}
