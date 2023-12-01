package config

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"golang.org/x/time/rate"
)

type Config struct {
	Port                   string        `json:"port"`
	Host                   string        `json:"host"`
	RedisAddr              string        `json:"redis-addr"`
	RedisPassword          string        `json:"redis-password"`
	RedisDB                int           `json:"redis-db"`
	ClientThresholdMinutes time.Duration `json:"client-threshold-minutes"`
	RateLimit              rate.Limit    `json:"rate-limit"`
	Burst                  int           `json:"burst"`
}

var ShurlConfig Config

func init() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config %v", err)
	}
	ShurlConfig = config
}

func loadConfig(filePath string) (Config, error) {
	var config Config
	//filePath = filepath.Join("..", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func GetConfig() Config {

	return ShurlConfig
}
