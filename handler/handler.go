package handler

import (
	"fmt"
	"net/http"
	"strings"
	"url-shortener/config"
	"url-shortener/shortener"
	"url-shortener/store"

	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//clean url
	if !strings.HasPrefix(creationRequest.LongUrl, "https://") && !strings.HasPrefix(creationRequest.LongUrl, "http://") {
		creationRequest.LongUrl = "https://" + creationRequest.LongUrl
	}
	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)
	//host := "http://localhost:8080/"
	host := config.GetConfig().Host
	url := host + shortUrl

	c.JSON(200, gin.H{
		"message":   "short url created successfuly",
		"short_url": url,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveOriginalurl(shortUrl)
	fmt.Println(shortUrl)
	c.Redirect(302, initialUrl)
}
