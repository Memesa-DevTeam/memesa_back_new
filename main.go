package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/nitishm/go-rejson/v4"
	"go.uber.org/fx"
	"memesa_go_backend/Config"
	"memesa_go_backend/Gin"
	"memesa_go_backend/MySQL"
	"memesa_go_backend/Redis"
)

func main() {
	fx.New(
		fx.Provide(
			Config.InitializeViperConfig,
			MySQL.InitializeMySQLDb,
			Redis.InitializeRedisDb,
			Gin.InitializeGinFramework,
		),
		// Invokes are here
		fx.Invoke(func(cfgCluster *Config.ConfigCluster) {}),
		fx.Invoke(func(SQLDB *sql.DB) {}),
		fx.Invoke(func(RedisDB *rejson.Handler) {}),
		fx.Invoke(func(GinFramework *gin.Engine) {}),
	).Run()
}
