package zaim

import (
	"fmt"
	"net/http"

	"github.com/dghubble/oauth1"
	"github.com/redcowe/go-zaim/config"
)

type Zaim struct {
	config *config.Config
	client *http.Client
	token  *oauth1.Token
}

func NewZaim() (*Zaim, error) {
	cfg, err := config.NewConfig()

	if err != nil {
		return nil, err
	}
	oauthcfg := oauth1.NewConfig(cfg.Credentials.ConsumerKey, cfg.Credentials.ConsumerSecret)
	token := oauth1.NewToken(cfg.Credentials.AccessToken, cfg.Credentials.AccessSecret)

	return &Zaim{
		config: cfg,
		token:  token,
		client: oauth1.NewClient(oauth1.NoContext, oauthcfg, token),
	}, nil
}

func (z *Zaim) VerifyAuthentication() bool {
	b, _ := z.client.Get("https://api.zaim.net/v2/home/user/verify")
	fmt.Println(b.StatusCode)
	return true
}
