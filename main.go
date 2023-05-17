package main

import (
	"database/sql"
	"go.uber.org/fx"
	"memesa_go_backend/Config"
	"memesa_go_backend/MySQL"
)

func main() {
	fx.New(
		fx.Provide(
			Config.InitializeViperConfig,
			MySQL.InitializeMySQLDb,
		),
		// Invokes are here
		fx.Invoke(func(cfgCluster *Config.ConfigCluster) {}),
		fx.Invoke(func(SQLDB *sql.DB) {}),
	).Run()
}
