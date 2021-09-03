package config

import "os"

type DBConfig struct {
	User     string
	Password string
	Driver   string
	Name     string
	Host     string
	Port     string
	SslMode  string
	Timezone string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
		Timezone:	os.Getenv("DB_TIMEZONE"),
	}
}
