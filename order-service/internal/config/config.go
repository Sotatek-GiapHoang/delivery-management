package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	ServerAddress string
	LogLevel      string
	JWTSecretKey  string
	KafkaBrokers  []string
	KafkaTopic    string
}

func LoadEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found. Using system environment variables.")
	}

	config := &Config{
		DatabaseURL:   buildDatabaseURL(),
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		LogLevel:      getEnv("LOG_LEVEL", "info"),
		JWTSecretKey:  getEnv("JWT_SECRET_KEY", ""),
		KafkaBrokers:  strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ","),
		KafkaTopic:    getEnv("KAFKA_TOPIC", "order_events"),
	}

	if config.JWTSecretKey == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY is not set")
	}

	return config, nil
}

func buildDatabaseURL() string {
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "user")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "delivery_db")
	port := getEnv("DB_PORT", "5432")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
