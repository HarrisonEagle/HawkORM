package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/HarrisonEagle/HawkORM/clauses"
	"github.com/HarrisonEagle/HawkORM/utils"
)

type MySQLDeleteClause struct {
	dbpool          *sql.DB
	Processor       *utils.Processor
	parser          *utils.Parser
	tableName       string
	whereConditions *WhereCondition
	orderBy         []string
	limit           int
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
	sc.whereConditions.SetAND(condition)
	return sc
}

func (sc *MySQLDeleteClause) WhereOr(condition interface{}) clauses.DeleteClause {
	sc.whereConditions.SetOR(condition)
	return sc
}

func (sc *MySQLDeleteClause) WhereNot(condition interface{}) clauses.DeleteClause {
	sc.whereConditions.SetNOT(condition)
	return sc
}

// DELETE FROM db.users WHERE cognito_id = "BBBBBB" ORDER BY name ASC LIMIT 2
func (sc *MySQLDeleteClause) GetSQLQuery() string {
	whereCond := sc.whereConditions.getConditionQuery()
	orderCond := ""
	limitCond := ""
	if len(sc.orderBy) > 0 {
		orderCond = "ORDER BY " + strings.Join(sc.orderBy, ", ")
	}
	if sc.limit > 0 {
		limitCond = fmt.Sprintf("LIMIT %d", sc.limit)
	}
	query := "DELETE FROM " + sc.tableName
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

func (sc *MySQLDeleteClause) Exec() (sql.Result, error) {
	return sc.Processor.ExecQuery(sc.GetSQLQuery())
}
