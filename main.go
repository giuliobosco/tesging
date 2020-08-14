package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	r := gin.Default()

	// register the ping endpoint
	r.GET("/ping", pingEndpoint)

	r.LoadHTMLGlob("templates/*")

	r.GET("/", indexEndpiont)

	return r
}

func pingEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func indexEndpiont(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)
}
