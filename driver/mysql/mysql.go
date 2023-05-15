package mysql

import (
	"database/sql"
	"github.com/HarrisonOwl/GABAORM/clauses"
	"github.com/HarrisonOwl/GABAORM/driver"
	"github.com/HarrisonOwl/GABAORM/utils"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	dbpool  *sql.DB
	parser  *utils.Parser
	scanner *utils.Scanner
}

func OpenMySQL(config string) (driver.Database, error) {
	db, err := sql.Open("mysql", config)
	if err != nil {
		return nil, err
	}
	dbController := &MySQLDB{dbpool: db, parser: &utils.Parser{}, scanner: utils.NewScanner(db)}
	return dbController, nil
}

func (db MySQLDB) Select(dataType interface{}) clauses.SelectClause {
	return &MySQLSelectClause{
		dbpool:     db.dbpool,
		scanner:    db.scanner,
		primaryKey: db.parser.ExtractPrimaryKey(dataType),
		tableName:  db.parser.GetTableName(dataType),
		columns:    db.parser.ExtractAllColumnsFromStruct(dataType, false),
		condition:  newCondition(db.parser),
	}
}
