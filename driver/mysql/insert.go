package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/HarrisonEagle/HawkORM/clauses"
	"github.com/HarrisonEagle/HawkORM/utils"
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

func (sc *MySQLInsertClause) GetSQLQuery() string {
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
	return sc.Processor.ExecQuery(sc.GetSQLQuery())
}
