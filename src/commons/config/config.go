package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config struct holds the configuration values.
type Config struct {
	EnvMode   string
	Port      string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	JwtSecret string
}

// LoadConfig loads the configuration values from the .env file.
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		EnvMode:   getEnv("GIN_MODE", "debug"), // "debug" / "release"
		Port:      getEnv("PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "127.0.0.1"),
		DBPort:    getEnv("DB_PORT", "3306"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASS", ""),
		DBName:    getEnv("DB_NAME", ""),
		JwtSecret: getEnv("JWT_SECRET", "secret"),
	}
}

// getEnv retrieves the value of the specified environment variable or uses a default value.
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
