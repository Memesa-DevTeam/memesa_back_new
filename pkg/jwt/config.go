package jwt

import (
	"fmt"
	"github.com/spf13/viper"
)

type JwtConfig struct {
	Issuer       string   `yaml:"issuer"`
	Secret       string   `yaml:"secret"`
	ExcludePaths []string `yaml:"exclude"`
}

func NewJwtConfig(v *viper.Viper) *JwtConfig {
	cfg := new(JwtConfig)
	if !v.IsSet("token") {
		return applyDefaultConfig()
	}
	if err := v.UnmarshalKey("token", cfg); err != nil {
		return applyDefaultConfig()
	}
	return cfg
}

func applyDefaultConfig() *JwtConfig {
	var emptyExclude []string
	cfg := &JwtConfig{
		Issuer:       "memesa-gin",
		Secret:       "memesa",
		ExcludePaths: emptyExclude,
	}
	fmt.Println("Using default Jwt Configuration...")
	return cfg
}
