package server

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Address   string `yaml:"address"`
	Port      int    `yaml:"port"`
	AllowCORS bool   `yaml:"allowCORS"`
}

func NewServerConfig(v *viper.Viper) *ServerConfig {
	cfg := new(ServerConfig)
	if !v.IsSet("server") {
		return applyDefaultConfig()
	}
	if err := v.UnmarshalKey("server", cfg); err != nil {
		return applyDefaultConfig()
	}
	return cfg
}

// applyDefaultConfig Returns the default config for server without reading from config file.
func applyDefaultConfig() *ServerConfig {
	cfg := &ServerConfig{
		Address:   "0.0.0.0",
		Port:      8080,
		AllowCORS: true,
	}
	fmt.Println("Using Default Server Configuration...")
	return cfg
}
