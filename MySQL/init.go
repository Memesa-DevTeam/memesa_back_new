package MySQL

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
	"memesa_go_backend/Config"
)

var SDB *sql.DB

// InitializeMySQLDb
// Initialize your MySQL Database and try the connection
func InitializeMySQLDb(lc fx.Lifecycle) *sql.DB {
	var err error
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Get the configuration and apply to database
			fmt.Println(fmt.Sprintf("%s:%s@tcp(%s:%s)/memesa", Config.CfgCluster.MySQL.Username, Config.CfgCluster.MySQL.Password, Config.CfgCluster.MySQL.Address, Config.CfgCluster.MySQL.Port))
			SDB, err = sql.Open("mysql", fmt.Sprint("%s:%s@tcp(%s:%s)/memesa", Config.CfgCluster.MySQL.Username, Config.CfgCluster.MySQL.Password, Config.CfgCluster.MySQL.Address, Config.CfgCluster.MySQL.Port))
			if err != nil {
				fmt.Println(err)
				panic("MySQL Database Initialization Error")
			}
			fmt.Println("MySQL Database has been configured successfully")
			return nil
		},
	})
	// Return the config database
	return SDB
}
