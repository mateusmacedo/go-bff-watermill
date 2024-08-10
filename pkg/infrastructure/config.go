package infrastructure

import (
	"os"
)

type Config struct {
	ServerPort      string
	MaxConnections  int
	EnableDebugMode bool
}

func LoadConfig() *Config {
	config := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	return config
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
