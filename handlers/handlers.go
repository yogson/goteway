package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"module-rest/logics"
	"net/http"
)

func init() {
	RegisterHandler(CreateNewEntity)
	RegisterHandler(GetTest)
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
