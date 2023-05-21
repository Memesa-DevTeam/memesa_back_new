package database

import (
	"fmt"
	"github.com/spf13/viper"
)

type SQLConfig struct {
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type RedisConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

// NewSQLConfig MySQL Config
func NewSQLConfig(v *viper.Viper) *SQLConfig {
	cfg := new(SQLConfig)
	if v.IsSet("mysql") {
		return applyDefaultSQLConfig()
	}
	if err := v.UnmarshalKey("mysql", cfg); err != nil {
		return applyDefaultSQLConfig()
	}
	fmt.Println("SQL Config Initialized successfully")
	return cfg
}

func applyDefaultSQLConfig() *SQLConfig {
	cfg := &SQLConfig{
		Address:  "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "",
	}
	return cfg
}

func NewRedisConfig(v *viper.Viper) *RedisConfig {
	cfg := new(RedisConfig)
	if !v.IsSet("redis") {
		return applyDefaultRedisConfig()
	}
	if err := v.UnmarshalKey("redis", cfg); err != nil {
		return applyDefaultRedisConfig()
	}
	return cfg
}

func applyDefaultRedisConfig() *RedisConfig {
	cfg := &RedisConfig{
		Address: "127.0.0.1",
		Port:    6379,
	}
	return cfg
}
