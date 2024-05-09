package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hey, gin URL shortener",
		})
	})

	err := r.Run(":1488")
	if err != nil {
		panic(fmt.Sprintf("Failed to start web-server - Error: %v", err))
	}
}
