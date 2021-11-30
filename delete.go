package GABAORM

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func(db *DB) Delete(models interface{},conditions string) error{
	valuePtr := reflect.ValueOf(models)
	value := valuePtr.Elem()
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array{
		if value.Len()== 0 && conditions == ""{
			return errors.New("Gloabal Delete Is Not Allowed!")
		}else if value.Len() == 0 {
			typeinf := value.Type().Elem()
			tableName := getTableName(typeinf.Name())
			query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName,conditions)
			log.Println(query)
			_ , err := db.dbpool.Exec(query)
			if err != nil{
				log.Println(err)
				return err
			}
		}else{
			//TODO: IMPLEMENT DELETE FROM ARRAY
		}
	}else if value.Kind() == reflect.Struct{
		typeinf := reflect.TypeOf(models).Elem()
		tableName := getTableName(typeinf.Name())
		if conditions != ""{
			query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName,conditions)
			log.Println(query)
			_ , err := db.dbpool.Exec(query)
			if err != nil{
				log.Println(err)
				return err
			}
		}
	}
	return nil
}
