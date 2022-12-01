package sdk

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetName() string
	Get(params ...map[string]any) func(c *gin.Context)
}

type Named struct {
	Name string
}

func (h Named) GetName() string {
	return h.Name
}

type ParametrizedHandler struct {
	Named
	handlerFunc func(params map[string]any) func(c *gin.Context)
}

type SimpleHandler struct {
	Named
	HandlerFunc func(c *gin.Context)
}

func (h SimpleHandler) Get(_ ...map[string]any) func(c *gin.Context) {
	return h.HandlerFunc
}

func (h ParametrizedHandler) Get(params ...map[string]any) func(c *gin.Context) {
	return h.handlerFunc(params[0])
}
