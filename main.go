package main

import (
	"fmt"
	"url-shortener/config"
	"url-shortener/handler"
	"url-shortener/store"

	"github.com/gin-gonic/gin"
)

func init() {
	store.InitializeStore()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Url shortener running",
		})

	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})
	port := config.GetConfig().Port
	fmt.Println(config.GetConfig())
	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
