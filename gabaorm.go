package GABAORM

import (
	"github.com/HarrisonOwl/GABAORM/driver"
	"github.com/HarrisonOwl/GABAORM/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func OpenMySQL(config string) (driver.Database, error) {
	return mysql.OpenMySQL(config)
}
