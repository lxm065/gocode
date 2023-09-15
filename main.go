package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	r := gin.Default()
	r.Use(favicon.New("./1.ico"))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world!",
		})
	})
	// r.Run() // listen and serve on 0.0.0.0:8080
	// r.Run(":9000")

	r.Run()
}
