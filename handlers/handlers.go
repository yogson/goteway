package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yogson/goteway/logics"
	"github.com/yogson/goteway/sdk"
	"github.com/yogson/goteway/server"
)

func init() {
	server.Register(sdk.SimpleHandler{sdk.Named{"simple"}, GetTest})
	server.Register(sdk.SimpleHandler{sdk.Named{"create-entity"}, CreateNewEntity})
}

func CreateNewEntity(c *gin.Context) {
	rawData, _ := c.GetRawData()
	var newEntity logics.Entity
	err := json.Unmarshal(rawData, &newEntity)
	if err != nil {
		return
	}
	newEntity.Save()
	c.String(http.StatusOK, newEntity.Serialize())
}

func GetTest(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
