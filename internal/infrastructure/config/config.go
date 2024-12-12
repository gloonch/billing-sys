package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    string
	HTTPPort   int
}

func LoadConfig() (*Config, error) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	httpPort, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     port,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SSLMode:    os.Getenv("DB_SSLMODE"),
		HTTPPort:   httpPort,
	}, nil
}
