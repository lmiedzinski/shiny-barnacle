package config

import "os"

type Config struct {
	PortNumber string
	BaseUrl    string
	LogLevel   string
}

func GetConfig() *Config {
	return &Config{
		PortNumber: getEnv("CFG_PORTNUMBER", "80"),
		BaseUrl:    getEnv("CFG_BASEURL", "/"),
		LogLevel:   getEnv("CFG_LOGLEVEL", "info"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
