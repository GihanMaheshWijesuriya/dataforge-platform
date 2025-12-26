package config

import "os"

type Config struct {
	Port string
}

func Load() Config {
	port := getEnv("PORT", "8081")
	return Config{Port: port}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
