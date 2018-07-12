// Package rcache is used to setup connection to REDIS
package rcache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

// InitCache establishes connection to redis in-memory cache.
func InitCache() (redis.Conn, error) {
	conStr := viper.GetString("services.cache.URL")
	return redis.Dial("tcp", conStr)
}
