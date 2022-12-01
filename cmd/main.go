package main

import (
	"github.com/yogson/goteway/server"
	"github.com/yogson/goteway/utils"
	_ "github.com/yogson/goteway/handlers"
	_ "github.com/yogson/goteway/middleware"
)

var version string

var configFile = utils.GetEnv("GW-CONFIG", "../api.yaml")

func main() {
	config := utils.ReadFile(configFile)
	srv := server.New().WithYamlConfig(config)
	srv.Run()
}
