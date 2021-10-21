package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World 2")
	})
	r.POST("/", func(c *gin.Context) {
		c.String(200, "Hello World 2")
	})
	err := r.Run(":8080")
	if err != nil {
		log.Panic(err)
	}
}
