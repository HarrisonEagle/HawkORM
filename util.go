package GABAORM

import (
	"strings"
	"unicode"
)

var layout = "2006-01-02 15:04:05"

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func getTableName(tablename string) string{
	chars := []rune(tablename)
	result := ""
	underline := false
	afterupper := false
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		if(IsUpper(char)){
			if underline && !afterupper{
				result += "_"
			}
			result += strings.ToLower(char)
			underline = true
			afterupper = true
		}else{
			result += char
			afterupper = false
		}
		if i == len(chars) -1 {
			if char == "s"{
				result += "es"
			}else  {
				result += "s"
			}
		}
	}
	return result
}

func getColumnName(columnName string) string{
	chars := []rune(columnName)
	result := ""
	underline := false
	afterupper := false
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		if IsUpper(char) {
			if underline && !afterupper{
				result += "_"
			}
			result += strings.ToLower(char)
			underline = true
			afterupper = true
		}else{
			result += char
			afterupper = false
		}
	}
	return result
}


