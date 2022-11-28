package handlers

import (
	"goteway/sdk"
	"goteway/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	server.Register(sdk.SimpleHandler{sdk.Named{"bomzh"}, GetBomzh})
}

func GetBomzh(c *gin.Context) {
	c.String(http.StatusOK, "You are BOMZH!")
}
