package mysql

import (
	"database/sql"
	"fmt"
	"github.com/HarrisonOwl/GABAORM/clauses"
	"github.com/HarrisonOwl/GABAORM/utils"
	"log"
	"strings"
)

type MySQLUpdateClause struct {
	dbpool    *sql.DB
	Processor *utils.Processor
	parser    *utils.Parser
	tableName string
	columns   []string
	values    []string
	condition *Condition
	orderBy   []string
	limit     int
}

func (sc *MySQLUpdateClause) SetData(data interface{}) clauses.UpdateClause {
	sc.columns = sc.parser.ExtractAllColumnsFromStructOrSlice(data, true)
	sc.values = sc.parser.ExtractAllValuesFromStruct(data, true)
	return sc
}

func (sc *MySQLUpdateClause) generateQuery() string {
	var valueConds []string
	whereCond := sc.condition.getConditionQuery()
	orderCond := ""
	limitCond := ""
	if len(sc.orderBy) > 0 {
		orderCond = "ORDER BY " + strings.Join(sc.orderBy, ", ")
	}
	if sc.limit > 0 {
		limitCond = fmt.Sprintf("LIMIT %d", sc.limit)
	}
	for i := 0; i < len(sc.values); i++ {
		valueConds = append(valueConds, fmt.Sprintf("%s = \"%s\"", sc.columns[i], sc.values[i]))
	}
	query := fmt.Sprintf("UPDATE %s SET %s %s %s %s", sc.tableName, strings.Join(valueConds, ", "), whereCond, orderCond, limitCond)
	log.Printf("Executing Query: %s \n", query)
	return query
}

func (sc *MySQLUpdateClause) Limit(number int) clauses.UpdateClause {
	sc.limit = number
	return sc
}

// TODO: string to struct type
func (sc *MySQLUpdateClause) OrderBy(orderBy []string) clauses.UpdateClause {
	sc.orderBy = orderBy
	return sc
}

func (sc *MySQLUpdateClause) Where(condition interface{}) clauses.UpdateClause {
	sc.condition.SetAND(condition)
	return sc
}

func (sc *MySQLUpdateClause) WhereOr(condition interface{}) clauses.UpdateClause {
	sc.condition.SetOR(condition)
	return sc
}

func (sc *MySQLUpdateClause) WhereNot(condition interface{}) clauses.UpdateClause {
	sc.condition.SetNOT(condition)
	return sc
}

func (sc *MySQLUpdateClause) Exec() (sql.Result, error) {
	return sc.Processor.ExecQuery(sc.generateQuery())
}
