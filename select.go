package GABAORM

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)


// Select all records when conditions is empty
func(db *DB) GetAll(models interface{},conditions string) error{
	valuePtr := reflect.ValueOf(models)
	value := valuePtr.Elem()
	if value.Kind() == reflect.Slice{
		typeinf := value.Type().Elem()
		tableName := getTableName(typeinf.Name())
		columns := ""
		count := 0
		extractColumnsFromStructWithCount(reflect.New(typeinf).Elem().Interface(),&columns,&count)
		query := fmt.Sprintf("SELECT %s FROM %s", columns, tableName)
		if conditions != ""{
			query += (" WHERE "+conditions)
		}
		log.Println(query)
		rows, err := db.dbpool.Query(query)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var columns []interface{}
			result := reflect.New(typeinf).Elem().Addr().Interface()
  			AssignFromArgs(result,&columns,0,count)
			err := rows.Scan(columns...)
			value.Set(reflect.Append(value, reflect.Indirect(reflect.ValueOf(result))))
			if err != nil {
				fmt.Println(err)
			}
		}
		return nil
	}else {
		return errors.New("Not Slice")
	}
}

func (db *DB) GetFirst()  {

}

func (db *DB) GetLast()  {

}
