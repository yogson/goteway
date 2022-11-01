package main

import (
	"github.com/gin-gonic/gin"
	"rest-playground/endpoints"
	"rest-playground/handlers"
	"strings"
)

var apiBase = "/api/v0.1/"

func getRouter() *gin.Engine {
	router := gin.New()
	var eps = endpoints.Endpoints{}
	eps.LoadFromYaml("endpoints.yaml")

	for _, ep := range eps.Endpoints {
		for _, methodHandler := range ep.Methods {
			handler := handlers.HandlerFunctions[methodHandler.Handler]
			router.Handle(
				strings.ToUpper(methodHandler.Method), apiBase+ep.Path, handler)
		}
	}
	return router
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := getRouter()
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
