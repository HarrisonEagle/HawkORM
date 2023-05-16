package mysql

import (
	"database/sql"
	"fmt"
	"github.com/HarrisonOwl/GABAORM/clauses"
	"github.com/HarrisonOwl/GABAORM/utils"
	"log"
	"strings"
)

type MySQLSelectClause struct {
	dbpool     *sql.DB
	Processor  *utils.Processor
	primaryKey string
	tableName  string
	columns    []string
	condition  *Condition
	orderBy    []string
	limit      int
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
	sc.condition.SetAND(condition)
	return sc
}

func (sc *MySQLSelectClause) WhereOr(condition interface{}) clauses.SelectClause {
	sc.condition.SetOR(condition)
	return sc
}

func (sc *MySQLSelectClause) WhereNot(condition interface{}) clauses.SelectClause {
	sc.condition.SetNOT(condition)
	return sc
}

func (sc *MySQLSelectClause) generateQuery() string {
	whereCond := sc.condition.getConditionQuery()
	orderCond := ""
	limitCond := ""
	columnCond := strings.Join(sc.columns, ", ")
	if len(sc.orderBy) > 0 {
		orderCond = "ORDER BY " + strings.Join(sc.orderBy, ", ")
	}
	if sc.limit > 0 {
		limitCond = fmt.Sprintf("LIMIT %d", sc.limit)
	}

	query := fmt.Sprintf("SELECT %s FROM %s %s %s %s", columnCond, sc.tableName, whereCond, orderCond, limitCond)
	log.Printf("Executing Query: %s \n", query)
	return query
}

func (sc *MySQLSelectClause) All(target interface{}) error {
	return sc.Processor.ScanQuery(target, sc.generateQuery())
}

func (sc *MySQLSelectClause) First(target interface{}) error {
	sc.Limit(1)
	sc.OrderBy([]string{fmt.Sprintf("%s ASC", sc.primaryKey)})
	return sc.Processor.ScanQuery(target, sc.generateQuery())
}

func (sc *MySQLSelectClause) Last(target interface{}) error {
	sc.Limit(1)
	sc.OrderBy([]string{fmt.Sprintf("%s DESC", sc.primaryKey)})
	return sc.Processor.ScanQuery(target, sc.generateQuery())
}
