package main

import (
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":     "pong",
		"description": "hello world",
	})

}

func main() {
	r := gin.Default()
	r.GET("/ping", pong)
	r.Run()
}
