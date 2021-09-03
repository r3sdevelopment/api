package application

import (
	"api/config"
	db "api/db"
	http "api/http"
)

func Start(cfg *config.Config) {
	db.ConnectDb(cfg)
	http.Start(cfg)
}
