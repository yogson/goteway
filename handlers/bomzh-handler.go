package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	RegisterHandler(GetBomzh)
}

func GetBomzh(c *gin.Context) {
	c.String(http.StatusOK, "You are BOMZH!")
}
