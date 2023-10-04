package HawkORM

import (
	"github.com/HarrisonEagle/HawkORM/driver"
	"github.com/HarrisonEagle/HawkORM/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func OpenMySQL(config string) (driver.Database, error) {
	return mysql.OpenMySQL(config)
}
