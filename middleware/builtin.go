package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/yogson/goteway/sdk"
	"github.com/yogson/goteway/server"
)

func init() {
	server.Register(sdk.SimpleHandler{sdk.Named{"trace"}, AddTrace})
}


func AddTrace(c *gin.Context) {
	// generate trace ID, set header
	traceParent := fmt.Sprintf("%s-%s-%s-%s", "00", "234ewedfasdfwefq3re", "234dfef3q4", "01")
	
	c.Header("traceparent", traceParent)
	c.Next()
}