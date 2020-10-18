package apiserver

import "github.com/ktroitskiy/http-rest-api/internal/store"

// Config ...
type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
	Store       *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
		Store:       store.NewConfig(),
	}
}
