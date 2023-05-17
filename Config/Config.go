package Config

import (
	"context"
	"fmt"
	_ "fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"os"
)

// structs definitions
type ConfigCluster struct {
	MySQL MySQL `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

type MySQL struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Redis struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type Token struct {
	Phrase string `yaml:"phrase"`
}

var CfgCluster *ConfigCluster
var mysqlConfig *MySQL
var redisConfig *Redis
var tokenConfig *Token

// readViperConfig
// Read and apply all configurations from viper
func readViperConfig() *ConfigCluster {
	// Read and apply working directory
	path, err := os.Getwd()
	if err != nil {
		panic("Unable to initialize working directory.")
	}
	// Get and apply configuration file
	vipConfig := viper.New()
	vipConfig.AddConfigPath(path + "/Config")
	vipConfig.SetConfigName("config")
	vipConfig.SetConfigType("yaml")
	// Read Configuration file
	if err := vipConfig.ReadInConfig(); err != nil {
		panic("Unable to read configuration file. Make sure the file is exist and accessible.")
	}
	// Convert to struct
	err = vipConfig.Unmarshal(&CfgCluster)
	if err != nil {
		panic("Convert Error")
	}
	// return
	return CfgCluster
}

// InitializeViperConfig
// Initialize configuration for server with lifecycle
func InitializeViperConfig(lc fx.Lifecycle) *ConfigCluster {
	var cfgCluster *ConfigCluster
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Read the configuration
			cfgCluster = readViperConfig()
			fmt.Println("Viper Configuration has been config")
			fmt.Println(cfgCluster.MySQL.Username)
			return nil
		},
	})
	return cfgCluster
}
