package GABAORM

import (
	"database/sql"
)

type DB struct {
	dbpool *sql.DB
}

func Open(driverName string, config string) (*DB,error){
	db, err := sql.Open(driverName, config)
	if err != nil {
		return nil , err
	}
	dbcontroller := &DB{dbpool: db}
	return dbcontroller,nil
}
