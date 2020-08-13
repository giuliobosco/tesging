package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// register the ping endpoint
	r.GET("/ping", pingEndpoint)

	r.Run()
}

func pingEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
