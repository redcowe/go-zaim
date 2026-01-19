package config

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
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

	credentials, err := loadCredentials()
	if err != nil {
		return nil, err
	}
	return &Config{
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
