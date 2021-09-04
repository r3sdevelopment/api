package keycloak

import (
	"api/config"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Keycloak struct {
	Client      *resty.Client
	Realm       string
	JwksUrl     string
	UserInfoUrl string
}

func New(c *config.Config) *Keycloak {
	baseUrl := c.Keycloak.URL
	client := resty.New()
	client.SetHostURL(baseUrl)

	realm := c.Keycloak.Realm
	userInfoUrl := fmt.Sprintf("%s/auth/realms/%s/protocol/openid-connect/userinfo", baseUrl, realm)
	jwksUrl := fmt.Sprintf("%s/auth/realms/%s/protocol/openid-connect/certs", baseUrl, realm)

	return &Keycloak{
		Client:      client,
		Realm:       realm,
		JwksUrl:     jwksUrl,
		UserInfoUrl: userInfoUrl,
	}
}

func (k *Keycloak) GetUserInfo(token string) {

	res, err := k.Client.R().SetAuthToken(token).Get(k.UserInfoUrl)
	if err != nil {
		fmt.Printf("Invalid token %s", err)
	}

	fmt.Printf("UserInfo %v\n", res)
}
