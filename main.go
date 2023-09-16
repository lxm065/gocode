package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/thinkerou/favicon"
)

func myHandler() gin.HandlerFunc {

	return func(context *gin.Context) {
		context.Set("usersession", "userid-1")
		context.Next()
		// context.Abort()
	}
}

func main() {

	//创建一个服务
	r := gin.Default()
	// r.Use(favicon.New("./1.ico"))

	//加载静态页面
	r.LoadHTMLGlob("./templates/*")

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

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"message": "这是go后台传来的数据",
		})
	})

	r.GET("/user/info", myHandler(), func(context *gin.Context) {
		//取出中间件中的值
		usersession := context.MustGet("usersession").(string)
		log.Println("++++++++++", usersession, "++++++++++")

		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})

	})

	// r.Run() // listen and serve on 0.0.0.0:8080
	r.Run(":9001")

	// r.Run()
}
