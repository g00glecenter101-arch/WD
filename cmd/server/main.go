package config

import (
	"os"
)

type Config struct {
	Port string
	// other config fields...
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default fallback
	}

	return &Config{
		Port: port,
	}
}
