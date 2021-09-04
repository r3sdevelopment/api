package config

import "os"

type HTTPConfig struct {
	IP   string
	Port string
	Env  string
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		IP: os.Getenv("IP"),
		Port: os.Getenv("PORT"),
		Env: os.Getenv("ENV"),
	}
}
