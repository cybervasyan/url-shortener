package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"url-shortener/handler"
	"url-shortener/store"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:1337"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:1338" || origin == "http://localhost:1488"
		},
		MaxAge: 12 * time.Hour,
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
