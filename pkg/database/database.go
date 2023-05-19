package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"
	"go.uber.org/fx"
	"log"
)

type Database struct {
	sql   *sql.DB
	redis *rejson.Handler
}

type InitialForm struct {
	VideoLikeDb    interface{} `json:"videoLikeDb"`
	CommentsLikeDb interface{} `json:"commentsLikeDb"`
	MomentsLikeDb  interface{} `json:"momentsLikeDb"`
}

func (d *Database) CreateRedisForm() error {
	InitializeRedisStructure(d.redis)
	return nil
}

func (d *Database) Backup() error {
	fmt.Println("Backing up...")
	return nil
}
func InitializeRedisStructure(d *rejson.Handler) {
	// Create the forms that the database will need
	var occupationObject = make(map[string]interface{}, 1)
	initForm := InitialForm{
		VideoLikeDb:    occupationObject,
		CommentsLikeDb: occupationObject,
		MomentsLikeDb:  occupationObject,
	}
	res, err := d.JSONSet("MEMESA_DB", "$", initForm)
	if err != nil {
		fmt.Println(err)
		log.Panicln("[Database/Redis.Initialization] Unable to initialize database structure")
	}
	if res.(string) == "OK" {
		fmt.Println("[Database/Redis.Initialization] Database Structure Build Success")
	}
}

// TODO: Unified Two New Functions
// Initialization Scripts
func initializeSQLDb(cfg *SQLConfig) *sql.DB {
	Db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/memesa", cfg.User, cfg.Password, cfg.Address, cfg.Port))
	if err != nil {
		panic(fmt.Sprintf("Unable to initialize SQL: %s", err))
		return nil
	}
	fmt.Println("SQL initialized successfully.")
	return Db
}

func initializeRedisDb(cfg *RedisConfig) *rejson.Handler {
	rh := rejson.NewReJSONHandler()
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.Address, cfg.Port))
	if err != nil {
		panic(fmt.Sprintf("Unable to initialize Redis: %s", err))
	}
	_, err = conn.Do("FLUSHALL")
	// Setup client
	rh.SetRedigoClient(conn)
	_, testErr := rh.JSONSet("Test", "$", "Hello")
	if testErr != nil {
		panic(fmt.Sprintf("Unable to connect to Redis: %s", testErr))
	}
	return rh
}

func NewDatabase(sqlConfig *SQLConfig, rConfig *RedisConfig) *Database {
	return &Database{
		sql:   initializeSQLDb(sqlConfig),
		redis: initializeRedisDb(rConfig),
	}
}

// lc
func lc(lifecycle fx.Lifecycle, db *Database) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.CreateRedisForm()
		},
		OnStop: func(ctx context.Context) error {
			return db.Backup()
		},
	})
}

func Provide() fx.Option {
	return fx.Options(fx.Provide(NewSQLConfig, NewRedisConfig, NewDatabase), fx.Invoke(lc))
}
