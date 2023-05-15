package utils

import (
	"database/sql"
	"errors"
	"reflect"
)

type Scanner struct {
	db *sql.DB
}

func NewScanner(db *sql.DB) *Scanner {
	return &Scanner{
		db: db,
	}
}

func (s *Scanner) ScanQuery(target interface{}, query string) error {
	valuePtr := reflect.ValueOf(target)
	value := valuePtr.Elem()
	var typeinf reflect.Type
	if value.Kind() == reflect.Slice {
		typeinf = value.Type().Elem()
	} else if value.Kind() == reflect.Struct {
		typeinf = value.Type()
	} else {
		errors.New("Error! Only Sturct or Slice can be used to scan!")
	}
	rows, err := s.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var columns []interface{}
		result := reflect.New(typeinf).Elem().Addr().Interface()
		s.assignFromArgs(result, &columns)
		err := rows.Scan(columns...)
		if value.Kind() == reflect.Slice {
			value.Set(reflect.Append(value, reflect.Indirect(reflect.ValueOf(result))))
		} else if value.Kind() == reflect.Struct {
			value.Set(reflect.Indirect(reflect.ValueOf(result)))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Scanner) assignFromArgs(model interface{}, columns *[]interface{}) {
	valueinf := reflect.ValueOf(model).Elem()
	typeinf := valueinf.Type()
	for i := 0; i < valueinf.NumField(); i++ {
		fieldkind := typeinf.Field(i).Type.Kind()
		if fieldkind == reflect.Slice || (fieldkind == reflect.Struct && typeinf.Field(i).Tag.Get("ForeignKey") != "") {
			continue
		}
		if typeinf.Field(i).Type.String() == "time.Time" {
			*columns = append(*columns, valueinf.Field(i).Addr().Interface())
		} else if fieldkind == reflect.Struct && typeinf.Field(i).Tag.Get("ForeignKey") == "" {
			s.assignFromArgs(valueinf.FieldByName(typeinf.Field(i).Name).Addr().Interface(), columns)
		} else {
			*columns = append(*columns, valueinf.Field(i).Addr().Interface())
		}
	}
}
