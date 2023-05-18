package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/fx"
)

type JwtClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func (c *JwtClaims) GenerateToken(cfg *JwtConfig) (string, error) {
	// method goes here
	return "nil", nil
}

func (c *JwtClaims) ParseToken(cfg *JwtConfig) (JwtClaims, error) {
	// method goes here
	return JwtClaims{}, nil
}

func Provide() fx.Option {
	return fx.Options(fx.Provide(NewJwtConfig))
}
