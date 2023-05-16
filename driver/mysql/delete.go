package mysql

import (
	"database/sql"
	"fmt"
	"github.com/HarrisonOwl/GABAORM/clauses"
	"github.com/HarrisonOwl/GABAORM/utils"
	"log"
	"strings"
)

type MySQLDeleteClause struct {
	dbpool    *sql.DB
	Processor *utils.Processor
	parser    *utils.Parser
	tableName string
	condition *Condition
	orderBy   []string
	limit     int
}

func (sc *MySQLDeleteClause) Limit(number int) clauses.DeleteClause {
	sc.limit = number
	return sc
}

// TODO: string to struct type
func (sc *MySQLDeleteClause) OrderBy(orderBy []string) clauses.DeleteClause {
	sc.orderBy = orderBy
	return sc
}

func (sc *MySQLDeleteClause) Where(condition interface{}) clauses.DeleteClause {
	sc.condition.SetAND(condition)
	return sc
}

func (sc *MySQLDeleteClause) WhereOr(condition interface{}) clauses.DeleteClause {
	sc.condition.SetOR(condition)
	return sc
}

func (sc *MySQLDeleteClause) WhereNot(condition interface{}) clauses.DeleteClause {
	sc.condition.SetNOT(condition)
	return sc
}

// DELETE FROM db.users WHERE cognito_id = "BBBBBB" ORDER BY name ASC LIMIT 2
func (sc *MySQLDeleteClause) generateQuery() string {
	whereCond := sc.condition.getConditionQuery()
	orderCond := ""
	limitCond := ""
	if len(sc.orderBy) > 0 {
		orderCond = "ORDER BY " + strings.Join(sc.orderBy, ", ")
	}
	if sc.limit > 0 {
		limitCond = fmt.Sprintf("LIMIT %d", sc.limit)
	}

	query := fmt.Sprintf("DELETE FROM %s %s %s %s", sc.tableName, whereCond, orderCond, limitCond)
	log.Printf("Executing Query: %s \n", query)
	return query
}

func (sc *MySQLDeleteClause) Exec() (sql.Result, error) {
	return sc.Processor.ExecQuery(sc.generateQuery())
}
