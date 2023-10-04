package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/HarrisonEagle/HawkORM/clauses"
	"github.com/HarrisonEagle/HawkORM/utils"
)

type MySQLUpdateClause struct {
	dbpool          *sql.DB
	Processor       *utils.Processor
	parser          *utils.Parser
	tableName       string
	columns         []string
	values          []string
	whereConditions *WhereCondition
	orderBy         []string
	limit           int
}

func (sc *MySQLUpdateClause) SetData(data interface{}) clauses.UpdateClause {
	sc.columns = sc.parser.ExtractAllColumnsFromStructOrSlice(data, true)
	sc.values = sc.parser.ExtractAllValuesFromStruct(data, true)
	return sc
}

func (sc *MySQLUpdateClause) GetSQLQuery() string {
	var valueConds []string
	whereCond := sc.whereConditions.getConditionQuery()
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
	query := fmt.Sprintf("UPDATE %s SET %s", sc.tableName, strings.Join(valueConds, ", "))
	if whereCond != "" {
		query += (" " + whereCond)
	}
	if orderCond != "" {
		query += (" " + orderCond)
	}
	if limitCond != "" {
		query += (" " + limitCond)
	}
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
	sc.whereConditions.SetAND(condition)
	return sc
}

func (sc *MySQLUpdateClause) WhereOr(condition interface{}) clauses.UpdateClause {
	sc.whereConditions.SetOR(condition)
	return sc
}

func (sc *MySQLUpdateClause) WhereNot(condition interface{}) clauses.UpdateClause {
	sc.whereConditions.SetNOT(condition)
	return sc
}

func (sc *MySQLUpdateClause) Exec() (sql.Result, error) {
	return sc.Processor.ExecQuery(sc.GetSQLQuery())
}
