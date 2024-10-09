package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	apiServer := gin.New()
	apiServer.LoadHTMLGlob("web/*")
	apiServer.GET("/hello", HelloWorldGet)
	apiServer.POST("/hello", HelloWorldPost)

	apiServer.Run(":3388")
}

func HelloWorldGet(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"username": "",
	})
}

func HelloWorldPost(context *gin.Context) {
	username := context.PostForm("UserName")
	context.HTML(http.StatusOK, "index.html", gin.H{
		"username": username,
	})
}
