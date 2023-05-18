package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

const (
	FILE_DIRECTORY = "./pkg/config"
	CONFIG_NAME    = "config"
	CONFIG_TYPE    = "yaml"
)

func NewViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName(CONFIG_NAME)
	v.SetConfigType(CONFIG_TYPE)
	v.AddConfigPath(FILE_DIRECTORY)
	// try to read the config
	if err := v.ReadInConfig(); err != nil {
		panic("Unable to read config file.")
	}
	fmt.Println("Viper initialized.")
	return v
}

func Provide() fx.Option {
	return fx.Options(fx.Provide(NewViper), fx.Invoke(lc))
}

func lc(lifecycle fx.Lifecycle, v *viper.Viper) {
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return v.WriteConfig()
		},
	})
}
