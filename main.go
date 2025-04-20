package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	App struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"app"`
}

func loadConfig() (*Config, error) {
	v := viper.New()

	// Set config paths
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")

	// Read environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env configuration: %v", err)
	}

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		log.Printf("Warning: failed to read config file: %v", err)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"config.app.name": config.App.Name,
			"env.app.name":    os.Getenv("APP_NAME"),
		})
	})

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}
