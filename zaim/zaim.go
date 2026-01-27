package zaim

import (
	"github.com/dghubble/oauth1"
	"github.com/redcowe/go-zaim/internal/config"
)

// Zaim is the main client for the Zaim API.
type Zaim struct {
	Money *MoneyService
	User  *UserService
}

// NewZaim creates a new Zaim client with configured sub-services.
func NewZaim() (*Zaim, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	oauthcfg := oauth1.NewConfig(cfg.Credentials.ConsumerKey, cfg.Credentials.ConsumerSecret)
	token := oauth1.NewToken(cfg.Credentials.AccessToken, cfg.Credentials.AccessSecret)
	client := oauth1.NewClient(oauth1.NoContext, oauthcfg, token)

	return &Zaim{
		Money: &MoneyService{client: client},
		User:  &UserService{client: client},
	}, nil
}
