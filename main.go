package main

import (
	"database/sql"
	"github.com/nitishm/go-rejson/v4"
	"go.uber.org/fx"
	"memesa_go_backend/Config"
	"memesa_go_backend/MySQL"
	"memesa_go_backend/Redis"
)

func main() {
	fx.New(
		fx.Provide(
			Config.InitializeViperConfig,
			MySQL.InitializeMySQLDb,
			Redis.InitializeRedisDb,
		),
		// Invokes are here
		fx.Invoke(func(cfgCluster *Config.ConfigCluster) {}),
		fx.Invoke(func(SQLDB *sql.DB) {}),
		fx.Invoke(func(RedisDB *rejson.Handler) {}),
	).Run()
}
