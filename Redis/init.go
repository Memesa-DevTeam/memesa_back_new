package Redis

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"
	"go.uber.org/fx"
	"memesa_go_backend/Config"
)

type initForm struct {
	VideoLikeDB    interface{} `json:"videoLikeDb"`
	CommentsLikeDB interface{} `json:"commentsLikeDb"`
	MomentsLikeDB  interface{} `json:"momentsLikeDb"`
}

var RDB = rejson.NewReJSONHandler()

// InitRedisDbStructure
// Fill in the basic structure for the Redis Database
func InitRedisDbStructure() {
	// Initialize database structure
	var occuObj = make(map[string]interface{}, 1)
	occuForm := initForm{
		VideoLikeDB:    occuObj,
		CommentsLikeDB: occuObj,
		MomentsLikeDB:  occuObj,
	}
	res, err := RDB.JSONSet("MEMESA_DB", "$", occuForm)
	if err != nil {
		panic("Unable to initialize Redis Database")
	}
	if res.(string) != "OK" {
		panic("Unable to initialize Redis Database. An error occured while initializing.")
	}
}

// InitializeRedisDb
// Initialize your Redis Database and try to connect
func InitializeRedisDb(lc fx.Lifecycle) *rejson.Handler {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", Config.CfgCluster.Redis.Address, Config.CfgCluster.Redis.Port))
			if err != nil {
				panic("Unable to connect to Redis")
			}
			// Clear the database
			_, err = conn.Do("FLUSHALL")
			// Setup trigger
			RDB.SetRedigoClient(conn)
			// Test connection
			_, testErr := RDB.JSONSet("Test", "$", "hello")
			if testErr != nil {
				panic("Unable to connect to Redis via RDB.")
			}
			// Initialize Structure
			InitRedisDbStructure()
			fmt.Println("Redis Database Has been initialized.")
			return nil
		},
	})
	return RDB
}
