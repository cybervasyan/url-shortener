package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"url-shortener/handler"
	"url-shortener/store"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:1337"}, // Ваш клиентский URL
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(context *gin.Context) {
		handler.CreateShortUrl(context)
	})

	r.GET("/:shortUrl", func(context *gin.Context) {
		handler.HandleShortUrlRedirect(context)
	})

	store.InitializeStore()

	err := r.Run(":1488")
	if err != nil {
		panic(fmt.Sprintf("Failed to start web-server - Error: %v", err))
	}
}
