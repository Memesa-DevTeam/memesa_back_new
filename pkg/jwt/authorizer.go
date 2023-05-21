package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"time"
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
	// set valid time
	expireTime := time.Now()
	if rememberMe {
		expireTime = time.Now().Add(24 * 365 * time.Hour)
	} else {
		expireTime = time.Now().Add(1 * time.Hour)
	}

	// generate claims user
	userClaims := TokenClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    a.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	// sign
	token, err := tokenClaims.SignedString(a.Secret)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return token, nil
}

func (a *Authorizer) ParseToken(token string) (*TokenClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.Secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*TokenClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func (a *Authorizer) CheckIsValid(token string) bool {
	tokenClaims, _ := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.Secret, nil
	})
	if tokenClaims == nil {
		return false
	}
	return tokenClaims.Valid
}
