package handlers

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime"
	"strings"
)

var HandlerFunctions = map[string]gin.HandlerFunc{}

func GetFunctionName(i interface{}) string {
	funcName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	funcNameExtra := strings.SplitAfter(funcName, ".")[0]
	return strings.Replace(funcName, funcNameExtra, "", 1)
}

func RegisterHandler(handler func(c *gin.Context)) {
	name := GetFunctionName(handler)
	HandlerFunctions[name] = handler
}
