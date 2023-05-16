package mysql

import (
	"database/sql"
	"fmt"
	"github.com/HarrisonOwl/GABAORM/clauses"
	"github.com/HarrisonOwl/GABAORM/utils"
	"log"
	"strings"
)

type MySQLInsertClause struct {
	dbpool    *sql.DB
	Processor *utils.Processor
	parser    *utils.Parser
	tableName string
	columns   []string
	values    [][]string
}

func (sc *MySQLInsertClause) SetData(data interface{}) clauses.InsertClause {
	sc.columns = sc.parser.ExtractAllColumnsFromStructOrSlice(data, true)
	sc.values = sc.parser.ExtractAllValuesFromStructOrSlice(data, true)
	return sc
}

func (sc *MySQLInsertClause) generateQuery() string {
	columnCond := "(" + strings.Join(sc.columns, ", ") + ")"
	var valueConds []string
	for i := 0; i < len(sc.values); i++ {
		var valueWithQuotes []string
		for j := 0; j < len(sc.values[i]); j++ {
			valueWithQuotes = append(valueWithQuotes, fmt.Sprintf("\"%s\"", sc.values[i][j]))
		}
		valueConds = append(valueConds, "("+strings.Join(valueWithQuotes, ", ")+")")
	}
	query := fmt.Sprintf("INSERT INTO %s %s VALUES %s", sc.tableName, columnCond, strings.Join(valueConds, ", "))
	log.Printf("Executing Query: %s \n", query)
	return query
}

func (sc *MySQLInsertClause) Exec() (sql.Result, error) {
	return sc.Processor.ExecQuery(sc.generateQuery())
}
