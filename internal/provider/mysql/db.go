// Package mysql is used to setup connection to MYSQL DB.
package mysql

import (
	"github.com/jmoiron/sqlx"
	// To load mysql driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// InitStore initializes DB connection with settings within config file.
func InitStore() (*sqlx.DB, error) {
	conStr := viper.GetString("services.db.URL")
	db, err := sqlx.Connect("mysql", conStr)
	return db, err
}
