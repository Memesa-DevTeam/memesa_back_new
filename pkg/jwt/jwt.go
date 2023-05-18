package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/fx"
	"time"
)

type JwtClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(userInfo JwtClaims, rememberMe bool, cfg *JwtConfig) (string, error) {
	// setup time
	expireTime := time.Now()
	if rememberMe {
		expireTime = expireTime.Add(24 * 365 * time.Hour)
	} else {
		expireTime = expireTime.Add(1 * time.Hour)
	}
	// Apply standard claims in to struct
	userInfo.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    cfg.Issuer,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, userInfo)
	token, err := tokenClaims.SignedString(cfg.Secret)
	if err != nil {
		fmt.Println(err)
	}
	// Validate
	if !CheckIsValid(token, cfg) {
		fmt.Println("invalidTokenValidation")
		return token, errors.New("invalidTokenValidation")
	}
	return token, nil
}

func ParseToken(token string, cfg *JwtConfig) (*JwtClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return cfg.Secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JwtClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func CheckIsValid(token string, cfg *JwtConfig) bool {
	tokenClaims, _ := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return cfg.Secret, nil
	})
	if tokenClaims == nil {
		return false
	}
	return tokenClaims.Valid
}

func Provide() fx.Option {
	return fx.Options(fx.Provide(NewJwtConfig))
}
