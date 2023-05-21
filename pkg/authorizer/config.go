package authorizer

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type AuthorizerConfig struct {
	ExcludePaths []string `yaml:"exclude"`
}

func NewAuthorizerConfig(v *viper.Viper) *AuthorizerConfig {
	cfg := new(AuthorizerConfig)
	if !v.IsSet("authorizer") {
		return applyDefaultConfig()
	}
	if err := v.UnmarshalKey("authorizer", cfg); err != nil {
		return applyDefaultConfig()
	}
	fmt.Println("Authorizer is initialized")
	return cfg
}

func applyDefaultConfig() *AuthorizerConfig {
	var initList = []string{""}
	cfg := &AuthorizerConfig{
		ExcludePaths: initList,
	}
	return cfg
}

func Provide() fx.Option {
	return fx.Options(fx.Provide(NewAuthorizerConfig))
}
