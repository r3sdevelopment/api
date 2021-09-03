package config

import "os"

type HTTPConfig struct {
	Port string
	Env  string
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Port: os.Getenv("PORT"),
		Env: os.Getenv("ENV"),
	}
}
