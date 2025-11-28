package config

import "os"

type Config struct {
	Address     string
	DatabaseURL string
	LogLevel    string
}

func Load() Config {
	return Config{
		Address:     getEnv("ADDRESS", ":8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@db:5432/orders?sslmode=disable"),
		LogLevel:    getEnv("LOGLEVEL", "INFO"),
	}
}

func getEnv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
