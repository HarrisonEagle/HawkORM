package GABAORM

import (
	"database/sql"
	"github.com/HarrisonOwl/GABAORM/driver"
	"github.com/HarrisonOwl/GABAORM/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	dbpool *sql.DB
}

func OpenMySQL(config string) (driver.Database, error) {
	return mysql.OpenMySQL(config)
}

func Open(driverName string, config string) (*DB, error) {
	db, err := sql.Open(driverName, config)
	if err != nil {
		return nil, err
	}
	dbcontroller := &DB{dbpool: db}
	return dbcontroller, nil
}
