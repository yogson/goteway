package handlers

import (
	"encoding/json"
	"goteway/logics"
	"goteway/sdk"
	"goteway/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	server.Register(sdk.SimpleHandler{sdk.Named{"simple"}, GetTest})
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
