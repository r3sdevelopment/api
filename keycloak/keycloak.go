package keycloak

import (
	"api/config"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Keycloak struct {
	Client *resty.Client
	realm  string
}

func New(cfg *config.Config) *Keycloak {
	client := resty.New()
	client.SetHostURL(cfg.Keycloak.URL)

	return &Keycloak{
		Client: client,
		realm:  cfg.Keycloak.Realm,
	}
}

func (keycloak *Keycloak) GetUserInfo(token string) {

	uri := fmt.Sprintf("/auth/realms/%s/protocol/openid-connect/userinfo", keycloak.realm)

	res, err := keycloak.Client.R().SetAuthToken(token).Get(uri)
	if err != nil {
		fmt.Sprintf("Invalid token %s", err)
	}

	fmt.Printf("UserInfo %v\n", res)
}
