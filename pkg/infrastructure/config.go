package infrastructure

import (
	"log"
	"os"
	"strconv"
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

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func init() {
	config := LoadConfig()
	log.Printf("Configuração carregada: %+v\n", config)
}
