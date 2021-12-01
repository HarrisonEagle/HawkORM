package GABAORM

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"
)

func(db *DB) Delete(models interface{},conditions string) error{
	valuePtr := reflect.ValueOf(models)
	value := valuePtr.Elem()
	var queries []QueryInf
	currenttime := time.Time{}
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array{
		typeinf := value.Type().Elem()
		tableName := getTableName(typeinf.Name())
		if value.Len()== 0 && conditions == ""{
			return errors.New("Gloabal Delete Is Not Allowed!")
		}else if value.Len() == 0 {
			query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName,conditions)
			log.Println("GABAORM Query:"+query)
			_ , err := db.dbpool.Exec(query)
			if err != nil{
				log.Println(err)
				return err
			}
		}else{
			//TODO: IMPLEMENT DELETE FROM ARRAY
			fmt.Println(value.Type())
			bulkModifyQuery(reflect.Indirect(value).Interface(),false,&queries,currenttime)
			query := fmt.Sprintf("DELETE FROM %s WHERE (%s) IN (%s)", tableName,queries[0].Columns,queries[0].Values)
			log.Println("GABAORM Query:"+query)
			_ , err := db.dbpool.Exec(query)
			if err != nil{
				log.Println(err)
				return err
			}
		}
	}else if value.Kind() == reflect.Struct{
		typeinf := reflect.TypeOf(models).Elem()
		tableName := getTableName(typeinf.Name())
		if conditions != ""{
			query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName,conditions)
			log.Println("GABAORM Query:"+query)
			_ , err := db.dbpool.Exec(query)
			if err != nil{
				log.Println(err)
				return err
			}
		}else{
			modifyQuery(reflect.Indirect(value).Interface(),false,&queries,currenttime)
			query := fmt.Sprintf("DELETE FROM %s WHERE (%s) IN (%s)", tableName,queries[0].Columns,queries[0].Values)
			log.Println("GABAORM Query:"+query)
			_ , err := db.dbpool.Exec(query)
			if err != nil{
				log.Println(err)
				return err
			}
		}
	}
	return nil
}
