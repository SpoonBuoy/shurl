package middleware

import (
	"fmt"
	"net/http"
	"time"
	"url-shortener/config"
	"url-shortener/store"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var clients *store.Clients
var cfg config.Config

func init() {
	clients = store.NewClientsTable()
	cfg = config.GetConfig()
	PerClientRateLimiterUpdater()
}

func PerClientRateLimiterUpdater() {

	go func() {
		for {
			time.Sleep(time.Minute)
			clients.Mu.Lock()
			for ip, client := range clients.Table {
				//clean user client from Client Table after
				//some threshold
				if time.Since(client.LastSeen) > cfg.ClientThresholdMinutes*time.Minute {
					delete(clients.Table, ip)
				}
			}
			clients.Mu.Unlock()
		}
	}()
}

func PerClientRateLimiter(c *gin.Context) {
	ip := c.ClientIP()
	fmt.Println("IP - : ", ip)
	clients.Mu.Lock()
	fmt.Println(cfg.RateLimit, cfg.Burst)
	if _, found := clients.Table[ip]; !found {
		clients.Table[ip] = &store.Client{Limiter: rate.NewLimiter(cfg.RateLimit, cfg.Burst)}
	}
	clients.Table[ip].LastSeen = time.Now()
	if !clients.Table[ip].Limiter.Allow() {
		clients.Mu.Unlock()
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "too many requests"})
		return
	}
	//fmt.Println((clients.Table[ip].Limiter))
	clients.Mu.Unlock()
	c.Next()
}
