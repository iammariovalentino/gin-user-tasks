package db

import (
	"fmt"
	"gin-user-tasks/src/pkg/config"
	"log"

	"github.com/jmoiron/sqlx"
)

type connDB struct {
	conn *sqlx.DB
}

func (c *connDB) GetConn() *sqlx.DB {
	return c.conn
}

func NewMySQL(config config.AppConfig) *connDB {
	descriptor := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&sql_mode=''&parseTime=true",
		config.DbConfig.DbUsername,
		config.DbConfig.DbPassword,
		config.DbConfig.DbHost,
		config.DbConfig.DbPort,
		config.DbConfig.DbName)

	db, err := sqlx.Connect("mysql", descriptor)
	if err != nil {
		log.Fatalf("error connecting to DB: %s", err)
	}

	db.SetMaxIdleConns(config.DbConfig.DbMaxIdleConnection)
	db.SetMaxOpenConns(config.DbConfig.DbMaxOpenConnection)

	return &connDB{conn: db}
}
