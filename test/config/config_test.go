package config

import (
	"os"
	"testing"

	"github.com/redcowe/go-zaim/config"
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
	os.Setenv("PORT", "8080")
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

func TestNewConfigLoadsPort(t *testing.T) {
	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("NewConfig() returned unexpected error: %v", err)
	}
	if cfg.Port != "8080" {
		t.Errorf("Port = %q, want %q", cfg.Port, "8080")
	}
}

func TestNewConfigDefaultPort(t *testing.T) {
	// Clear PORT to test default behavior
	os.Unsetenv("PORT")

	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("NewConfig() returned unexpected error: %v", err)
	}
	if cfg.Port != "8080" {
		t.Errorf("Port = %q, want default %q", cfg.Port, "8080")
	}

	// Restore PORT for other tests
	os.Setenv("PORT", "8080")
}

func TestNewConfigCustomPort(t *testing.T) {
	os.Setenv("PORT", "3000")

	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("NewConfig() returned unexpected error: %v", err)
	}
	if cfg.Port != "3000" {
		t.Errorf("Port = %q, want %q", cfg.Port, "3000")
	}

	// Restore PORT for other tests
	os.Setenv("PORT", "8080")
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
