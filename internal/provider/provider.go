package provider

import (
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	"github.com/govinda-attal/cabride-api/internal/provider/mysql"
	"github.com/govinda-attal/cabride-api/internal/provider/rcache"
)

const (
	// PrvDB is used as a key to store db connection within viper config data map.
	PrvDB = "prv.db"
	// PrvCache is used as a key to store cache connection within viper config data map.
	PrvCache = "prv.cache"
)

// Setup function loads providers for this microservice. For cabride-api mysql db is the only provider.
// This function is called at the microservice startup.
func Setup() {
	db, err := mysql.InitStore()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetDefault(PrvDB, db)
	rcon, err := rcache.InitCache()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetDefault(PrvCache, rcon)
}

// DB function returns sql DB connection for this microservice.
func DB() *sqlx.DB {
	return viper.Get(PrvDB).(*sqlx.DB)
}

// Cache function returns Redis connection for this microservice.
func Cache() redis.Conn {
	return viper.Get(PrvCache).(redis.Conn)
}

// Cleanup function cleans up active provider resources if any.
// This function is to be called when the microservice is shutting down.
func Cleanup() {
	if db := DB(); db != nil {
		db.Close()
	}
	if cache := Cache(); cache != nil {
		cache.Close()
	}
}
