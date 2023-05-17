package Config

import (
	"context"
	_ "fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// structs definitions
type ConfigCluster struct {
	MySQL MySQL `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

type MySQL struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

type Redis struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type Token struct {
	Phrase string `yaml:"phrase"`
}

// InitializeViperConfig
// Initialize all configurations while boot up
func InitializeViperConfig(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Initialize all info at here
			viper.SetConfigFile("./config.yaml")
		},
	})
}
