package config

import "os"

type KeycloakConfig struct {
	URL      string
	Realm    string
	ClientId string
}

func LoadKeycloakConfig() KeycloakConfig {
	return KeycloakConfig{
		URL:      os.Getenv("KEYCLOAK_URL"),
		Realm:    os.Getenv("KEYCLOAK_REALM"),
		ClientId: os.Getenv("KEYCLOAK_CLIENT_ID"),
	}
}
