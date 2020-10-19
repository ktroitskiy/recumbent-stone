package apiserver

// Config ...
type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
	}
}
