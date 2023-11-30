package handler

import (
	"fmt"
	"net/http"
	"strings"
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
	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)
	host := "http://localhost:8080/"
	url := host + shortUrl
	//clean url
	if !strings.HasPrefix(url, "https://") || !strings.HasPrefix(url, "http://") {
		url = "https://" + url
	}

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
