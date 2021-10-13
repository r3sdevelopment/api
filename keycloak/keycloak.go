package keycloak

import (
	"api/config"
	"api/pkg/entities"
	"api/utils"
	"fmt"
	"github.com/MicahParks/keyfunc"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

const RolesKey = "UserRoles"
const UserIdKey = "UserID"

type Keycloak struct {
	Client      *resty.Client
	Realm       string
	JwksUrl     string
	UserInfoUrl string
}

type Roles []string
type RealmRoles struct {
	Roles Roles
}

type ResourceRoles map[string]map[string][]string

// Address TODO what fields does any address have?

// Claims StringOrArray represents a value that can either be a string or an array of strings
type Claims struct {
	jwt.StandardClaims
	Aud           []string      `json:"aud,omitempty"`
	RealmRoles    RealmRoles    `json:"realm_access"`
	ResourceRoles ResourceRoles `json:"resource_access"`
	Roles         Roles         `json:"roles"`
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

func (k *Keycloak) ApplyMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqToken := c.Get(fiber.HeaderAuthorization)
		splitToken := strings.Split(reqToken, "Bearer ")

		if len(splitToken) == 2 {
			reqToken = splitToken[1]

			jwks, err := keyfunc.Get(k.JwksUrl)
			if err != nil {
				fmt.Printf("Failed to get the JWKs from the given URL. (Error: %s, URL: %s)", err.Error(), k.JwksUrl)

				return c.Next()
			}

			token, claimsErr := jwt.ParseWithClaims(reqToken, &Claims{}, jwks.KeyFunc)

			if claimsErr != nil {
				fmt.Printf("Failed to get parse claims. (Error: %s)", claimsErr.Error())

				return c.Next()
			}

			if claims, ok := token.Claims.(*Claims); ok && token.Valid {
				c.Locals(RolesKey, claims.Roles)
				c.Locals(UserIdKey, claims.Subject)
			}
		}

		return c.Next()
	}
}

func (k *Keycloak) NeedsRole(needsRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if roles, ok := c.Locals(RolesKey).(Roles); ok {
			for _, role := range needsRoles {
				if utils.Contains(roles, role) {
					return c.Next()
				}
			}
		}

		return c.Status(fiber.StatusUnauthorized).JSON(&entities.ApiResponse{
			Code:    fiber.StatusUnauthorized,
			Type:    "NotAuthorized",
			Message: "Not Authorized",
		})
	}
}
