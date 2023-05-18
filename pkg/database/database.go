package database

import (
	"database/sql"
	"fmt"
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

func InitializeRedisStructure(d *Database) {
	// Create the forms that the database will need
	var occupationObject = make(map[string]interface{}, 1)
	initForm := InitialForm{
		VideoLikeDb:    occupationObject,
		CommentsLikeDb: occupationObject,
		MomentsLikeDb:  occupationObject,
	}
	res, err := d.redis.JSONSet("MEMESA_DB", "$", initForm)
	if err != nil {
		fmt.Println(err)
		log.Panicln("[Database/Redis.Initialization] Unable to initialize database structure")
	}
	if res.(string) == "OK" {
		fmt.Println("[Database/Redis.Initialization] Database Structure Build Success")
	}
}

// Initialization Scripts
func NewSqlDb(d *Database, cfg *SQLConfig) *Database {
	Db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/memesa", cfg.User, cfg.Password, cfg.Address, cfg.Port))
	if err != nil {
		panic(fmt.Sprintf("Unable to initialize SQL: %s", err))
		return nil
	}
	d.sql = Db
	return d
}

func NewRedisDb(d *Database, cfg *RedisConfig) *Database {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.Address, cfg.Port))
	if err != nil {
		panic(fmt.Sprintf("Unable to initialize Redis: %s", err))
	}
	_, err = conn.Do("FLUSHALL")
	// Setup client
	d.redis.SetRedigoClient(conn)
	_, testErr := d.redis.JSONSet("Test", "$", "Hello")
	if testErr != nil {
		panic(fmt.Sprintf("Unable to connect to Redis: %s", testErr))
	}
	// Initialize Database Structure
	InitializeRedisStructure(d)
	return d
}

// lc

func Provide() fx.Option {
	return fx.Options(fx.Provide(NewSqlDb, NewRedisConfig))
}
