package config

import (
	"errors"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Credentials Credentials
}

type Credentials struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		slog.Info("port not found in env. defaulting to 8080")
		port = "8080"
	}

	credentials, err := loadCredentials()
	if err != nil {
		return nil, err
	}
	return &Config{
		Port:        port,
		Credentials: *credentials,
	}, nil
}

func loadCredentials() (*Credentials, error) {
	//first check the env
	consumerKey := strings.TrimSpace(os.Getenv("ZAIM_CONSUMER_KEY"))
	consumerSecret := strings.TrimSpace(os.Getenv("ZAIM_CONSUMER_SECRET"))
	accessToken := strings.TrimSpace(os.Getenv("ZAIM_ACCESS_TOKEN"))
	accessSecret := strings.TrimSpace(os.Getenv("ZAIM_ACCESS_SECRET"))

	if consumerKey != "" && consumerSecret != "" && accessToken != "" && accessSecret != "" {
		return &Credentials{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			AccessToken:    accessToken,
			AccessSecret:   accessSecret,
		}, nil
	}
	return nil, errors.New("unable to load credentials from environment")
}
