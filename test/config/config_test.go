package config

import (
	"os"
	"testing"

	"github.com/redcowe/go-zaim/internal/config"
)

func TestMain(m *testing.M) {

	setup()

	res := m.Run()

	os.Exit(res)
}

func setup() {
	os.Setenv("ZAIM_CONSUMER_KEY", "consumer_key")
	os.Setenv("ZAIM_CONSUMER_SECRET", "consumer_secret")
	os.Setenv("ZAIM_ACCESS_TOKEN", "access_token")
	os.Setenv("ZAIM_ACCESS_SECRET", "access_secret")
}

func TestGetNewConfigFromEnv(t *testing.T) {
	_, err := config.NewConfig()
	if err != nil {
		t.Fatalf("NewConfig() returned unexpected error: %v", err)
	}
}

func TestNewConfigLoadsAllCredentials(t *testing.T) {
	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("NewConfig() returned unexpected error: %v", err)
	}

	tests := []struct {
		name string
		got  string
		want string
	}{
		{"ConsumerKey", cfg.Credentials.ConsumerKey, "consumer_key"},
		{"ConsumerSecret", cfg.Credentials.ConsumerSecret, "consumer_secret"},
		{"AccessToken", cfg.Credentials.AccessToken, "access_token"},
		{"AccessSecret", cfg.Credentials.AccessSecret, "access_secret"},
	}

	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("%s = %q, want %q", tt.name, tt.got, tt.want)
		}
	}
}

func TestNewConfigMissingCredentials(t *testing.T) {
	// Clear a required credential
	os.Unsetenv("ZAIM_CONSUMER_KEY")

	_, err := config.NewConfig()
	if err == nil {
		t.Error("NewConfig() should return error when credentials are missing")
	}

	// Restore for other tests
	os.Setenv("ZAIM_CONSUMER_KEY", "consumer_key")
}
