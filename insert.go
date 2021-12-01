package GABAORM

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

func(db *DB) Insert(model interface{}) error{
	var queries []QueryInf
	currenttime := time.Now()
	if reflect.TypeOf(model).Kind() == reflect.Slice || reflect.TypeOf(model).Kind() == reflect.Array{
		if reflect.ValueOf(model).Len()== 0{
			return nil
		}
		bulkModifyQuery(model,false,&queries,currenttime)
	}else if reflect.TypeOf(model).Kind() == reflect.Struct{
		modifyQuery(model,false,&queries,currenttime)
	}
	tx, _ := db.dbpool.Begin()
	for i := len(queries) - 1;i >=0;i--{
		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", queries[i].TableName, queries[i].Columns,queries[i].Values)
		log.Println("GABAORM Query:"+query)
		_ , err := tx.Exec(query)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
