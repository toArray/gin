package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middlePrintInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Printf("paht:%v | method:%v\n", path, method)
		context.Next()
		fmt.Println(context.Writer.Status())
	}
}
