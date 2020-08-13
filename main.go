package main

import "github.com/gin-gonic/gin"

func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	r := gin.Default()

	// register the ping endpoint
	r.GET("/ping", pingEndpoint)

	return r
}

func pingEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
