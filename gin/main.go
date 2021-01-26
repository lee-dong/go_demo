package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {


	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "我是一个测试功能~")
	})
	r.Run(":8080")
}
