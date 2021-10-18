package application

import (
	"api/config"
	db "api/db"
	"api/http"
	"api/keycloak"
)

func Start(cfg *config.Config) {
	k := keycloak.New(cfg)

	db.Connect(cfg)
	http.Start(cfg, k)
}
