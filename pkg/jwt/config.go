package jwt

import (
	"fmt"
	"github.com/spf13/viper"
)

type JwtConfig struct {
	Issuer string `yaml:"issuer"`
	Secret string `yaml:"secret"`
}

func NewJwtConfig(v *viper.Viper) *JwtConfig {
	cfg := new(JwtConfig)
	if v.IsSet("token") {
		return applyDefaultConfig()
	}
	if err := v.UnmarshalKey("token", cfg); err != nil {
		return applyDefaultConfig()
	}
	fmt.Println("Jwt Service Initialized")
	return cfg
}

func applyDefaultConfig() *JwtConfig {
	cfg := &JwtConfig{
		Issuer: "memesa",
		Secret: "memesa",
	}
	fmt.Println("Using Default Jwt Configuration")
	return cfg
}
