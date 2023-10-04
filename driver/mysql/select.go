package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/HarrisonEagle/HawkORM/clauses"
	"github.com/HarrisonEagle/HawkORM/utils"
)

type MySQLSelectClause struct {
	dbpool          *sql.DB
	Processor       *utils.Processor
	primaryKey      string
	tableName       string
	columns         []string
	whereConditions *WhereCondition
	orderBy         []string
	limit           int
}

func (sc *MySQLSelectClause) Limit(number int) clauses.SelectClause {
	sc.limit = number
	return sc
}

// TODO: string to struct type
func (sc *MySQLSelectClause) OrderBy(orderBy []string) clauses.SelectClause {
	sc.orderBy = orderBy
	return sc
}

func (sc *MySQLSelectClause) Where(condition interface{}) clauses.SelectClause {
	sc.whereConditions.SetAND(condition)
	return sc
}

func (sc *MySQLSelectClause) WhereOr(condition interface{}) clauses.SelectClause {
	sc.whereConditions.SetOR(condition)
	return sc
}

func (sc *MySQLSelectClause) WhereNot(condition interface{}) clauses.SelectClause {
	sc.whereConditions.SetNOT(condition)
	return sc
}

func (sc *MySQLSelectClause) GetSQLQuery() string {
	whereCond := sc.whereConditions.getConditionQuery()
	orderCond := ""
	limitCond := ""
	columnCond := strings.Join(sc.columns, ", ")
	if len(sc.orderBy) > 0 {
		orderCond = "ORDER BY " + strings.Join(sc.orderBy, ", ")
	}
	if sc.limit > 0 {
		limitCond = fmt.Sprintf("LIMIT %d", sc.limit)
	}

	query := fmt.Sprintf("SELECT %s FROM %s", columnCond, sc.tableName)
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

func (sc *MySQLSelectClause) All(target interface{}) error {
	return sc.Processor.ScanQuery(target, sc.GetSQLQuery())
}

func (sc *MySQLSelectClause) First(target interface{}) error {
	sc.Limit(1)
	sc.OrderBy([]string{fmt.Sprintf("%s ASC", sc.primaryKey)})
	return sc.Processor.ScanQuery(target, sc.GetSQLQuery())
}

func (sc *MySQLSelectClause) Last(target interface{}) error {
	sc.Limit(1)
	sc.OrderBy([]string{fmt.Sprintf("%s DESC", sc.primaryKey)})
	return sc.Processor.ScanQuery(target, sc.GetSQLQuery())
}
