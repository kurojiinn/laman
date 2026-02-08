package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config содержит всю конфигурацию приложения.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Jaeger   JaegerConfig
}

// ServerConfig содержит конфигурацию сервера.
type ServerConfig struct {
	Port string
	Host string
}

// DatabaseConfig содержит конфигурацию базы данных.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// JWTConfig содержит конфигурацию JWT.
type JWTConfig struct {
	Secret string
}

// JaegerConfig содержит конфигурацию трейсинга Jaeger.
type JaegerConfig struct {
	Endpoint string
}

// Load загружает конфигурацию из переменных окружения.
func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "laman"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		},
		Jaeger: JaegerConfig{
			Endpoint: getEnv("JAEGER_ENDPOINT", "http://jaeger:14268/api/traces"),
		},
	}

	if cfg.JWT.Secret == "your-secret-key-change-in-production" {
		return nil, fmt.Errorf("JWT_SECRET должен быть установлен в переменных окружения")
	}

	return cfg, nil
}

// DSN возвращает строку подключения к базе данных.
func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
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
