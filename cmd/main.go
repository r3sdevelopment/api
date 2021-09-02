package main

import (
	application "api"
	"api/config"
)

func main() {
	cfg := config.NewConfig()

	application.Start(cfg)
}
