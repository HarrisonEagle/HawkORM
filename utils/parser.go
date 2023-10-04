package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Parser struct{}

func (p *Parser) ExtractPrimaryKey(model interface{}) string {
	var typeinf reflect.Type
	if reflect.ValueOf(model).Kind() == reflect.Ptr {
		typeinf = reflect.ValueOf(model).Elem().Type()
	} else {
		typeinf = reflect.TypeOf(model)
	}
	for i := 0; i < typeinf.NumField(); i++ {
		if typeinf.Field(i).Tag.Get("hawkorm") != "" {
			if strings.Contains(typeinf.Field(i).Tag.Get("hawkorm"), "primarykey") {
				return p.getColumnName(typeinf.Field(i).Name)
			}
		}
	}
	return ""
}

func (p *Parser) ExtractAllColumnsFromStructOrSlice(model interface{}, notZeroOnly bool) []string {
	var typeinf reflect.Type
	var valueinf reflect.Value
	if reflect.ValueOf(model).Kind() == reflect.Ptr {
		typeinf = reflect.ValueOf(model).Elem().Type()
		valueinf = reflect.ValueOf(model).Elem()
		if valueinf.Kind() == reflect.Slice || valueinf.Kind() == reflect.Array {
			typeinf = reflect.TypeOf(valueinf.Index(0).Interface())
			valueinf = valueinf.Index(0)
		}
	} else if reflect.ValueOf(model).Kind() == reflect.Slice || reflect.ValueOf(model).Kind() == reflect.Array {
		objects := reflect.ValueOf(model)
		typeinf = reflect.TypeOf(objects.Index(0).Interface())
		valueinf = objects.Index(0)
	} else {
		typeinf = reflect.TypeOf(model)
		valueinf = reflect.ValueOf(model)
	}
	var columns []string
	for i := 0; i < typeinf.NumField(); i++ {
		fieldType := typeinf.Field(i).Type.Kind()
		value := valueinf.Field(i)
		if notZeroOnly && value.IsZero() {
			continue
		}
		if fieldType == reflect.Struct && typeinf.Field(i).Type.String() == "time.Time" || fieldType == reflect.Int || fieldType == reflect.Float64 || fieldType == reflect.String {
			columns = append(columns, p.getColumnName(typeinf.Field(i).Name))
		}
	}
	return columns
}

func (p *Parser) ExtractAllValuesFromStructOrSlice(model interface{}, notZeroOnly bool) [][]string {
	var valueinf reflect.Value
	result := [][]string{}
	if reflect.ValueOf(model).Kind() == reflect.Ptr {
		valueinf = reflect.ValueOf(model).Elem()
		if valueinf.Kind() == reflect.Slice || valueinf.Kind() == reflect.Array {
			for i := 0; i < valueinf.Len(); i++ {
				value := valueinf.Index(i).Interface()
				fmt.Println(p.ExtractAllValuesFromStruct(value, notZeroOnly))
				result = append(result, p.ExtractAllValuesFromStruct(value, notZeroOnly))
			}
		} else {
			result = append(result, p.ExtractAllValuesFromStruct(valueinf.Interface(), notZeroOnly))
		}
	} else if reflect.ValueOf(model).Kind() == reflect.Slice || reflect.ValueOf(model).Kind() == reflect.Array {
		valueinf = reflect.ValueOf(model)
		for i := 0; i < valueinf.Len(); i++ {
			value := valueinf.Index(i).Interface()
			result = append(result, p.ExtractAllValuesFromStruct(value, notZeroOnly))
		}
	} else {
		result = append(result, p.ExtractAllValuesFromStruct(model, notZeroOnly))
	}
	return result
}

func (p *Parser) ExtractAllValuesFromStruct(model interface{}, notZeroOnly bool) []string {
	var typeinf reflect.Type
	var valueinf reflect.Value
	if reflect.ValueOf(model).Kind() == reflect.Ptr {
		typeinf = reflect.ValueOf(model).Elem().Type()
		valueinf = reflect.ValueOf(model).Elem()
	} else if reflect.ValueOf(model).Kind() == reflect.Struct {
		typeinf = reflect.TypeOf(model)
		valueinf = reflect.ValueOf(model)
	}
	var layout = "2006-01-02 15:04:05"
	var values []string
	for i := 0; i < typeinf.NumField(); i++ {
		fieldType := typeinf.Field(i).Type.Kind()
		value := valueinf.Field(i)
		if value.IsZero() {
			if notZeroOnly {
				continue
			} else {
				values = append(values, "null")
			}
		}
		/// TODO: support int8, int16, int32, int64, uint, float32
		if fieldType == reflect.Struct && typeinf.Field(i).Type.String() == "time.Time" {
			datetime := value.Interface().(time.Time)
			datestring := datetime.Format(layout)
			values = append(values, datestring)
		} else if fieldType == reflect.String {
			values = append(values, value.Interface().(string))
		} else if fieldType == reflect.Int {
			values = append(values, strconv.Itoa(value.Interface().(int)))
		} else if fieldType == reflect.Float64 {
			values = append(values, strconv.FormatFloat(value.Interface().(float64), 'f', -1, 64))
		}
	}
	return values
}

func (p *Parser) isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (p *Parser) GetTableName(dataType interface{}) string {
	var typeinf reflect.Type
	if reflect.ValueOf(dataType).Kind() == reflect.Ptr {
		typeinf = reflect.ValueOf(dataType).Elem().Type()
	} else {
		typeinf = reflect.TypeOf(dataType)
	}
	chars := []rune(typeinf.Name())
	result := ""
	underline := false
	afterupper := false
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		if p.isUpper(char) {
			if underline && !afterupper {
				result += "_"
			}
			result += strings.ToLower(char)
			underline = true
			afterupper = true
		} else {
			result += char
			afterupper = false
		}
		if i == len(chars)-1 {
			if char == "s" {
				result += "es"
			} else {
				result += "s"
			}
		}
	}
	return result
}

func (p *Parser) getColumnName(columnName string) string {
	chars := []rune(columnName)
	result := ""
	underline := false
	afterupper := false
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		if p.isUpper(char) {
			if underline && !afterupper {
				result += "_"
			}
			result += strings.ToLower(char)
			underline = true
			afterupper = true
		} else {
			result += char
			afterupper = false
		}
	}
	return result
}
