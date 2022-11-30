package handlers

import (
	"time"

	"goteway/sdk"
	"goteway/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	server.Register(sdk.SimpleHandler{sdk.Named{"datetime"}, GetDate})
}

func GetDate(c *gin.Context) {
	dt := time.Now()
	c.String(http.StatusOK, dt.String())
}
