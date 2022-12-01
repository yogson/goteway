package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yogson/goteway/sdk"
	"github.com/yogson/goteway/server"
)

func init() {
	server.Register(sdk.SimpleHandler{sdk.Named{"datetime"}, GetDate})
}

func GetDate(c *gin.Context) {
	dt := time.Now()
	c.String(http.StatusOK, dt.String())
}
