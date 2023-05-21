package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Authorizer struct {
	Issuer  string
	Secret  string
	Exclude []string
}

type TokenClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func Provide() fx.Option {
	return fx.Provide(NewAuthorizer, NewJwtConfig)
}

func NewAuthorizer(v *viper.Viper, config *JwtConfig) *Authorizer {
	auth := new(Authorizer)
	auth.Issuer = config.Issuer
	auth.Secret = config.Secret
	auth.Exclude = config.ExcludePaths
	if err := v.UnmarshalKey("token", &auth); err != nil {
		fmt.Println("Cannot initialize authorizer")
	}
	return auth
}

// Methods here
func (a *Authorizer) GenerateToken(username, password string, rememberMe bool) (string, error) {
	return "true", nil
}

func (a *Authorizer) ParseToken(token string) (*TokenClaims, error) {
	return nil, nil
}

func (a *Authorizer) CheckIsValid(token string) bool {
	return true
}
