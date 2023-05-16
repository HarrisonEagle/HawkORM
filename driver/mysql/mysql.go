package mysql

import (
	"database/sql"
	"github.com/HarrisonOwl/GABAORM/clauses"
	"github.com/HarrisonOwl/GABAORM/driver"
	"github.com/HarrisonOwl/GABAORM/utils"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	dbpool    *sql.DB
	parser    *utils.Parser
	Processor *utils.Processor
}

func OpenMySQL(config string) (driver.Database, error) {
	db, err := sql.Open("mysql", config)
	if err != nil {
		return nil, err
	}
	dbController := &MySQLDB{dbpool: db, parser: &utils.Parser{}, Processor: utils.NewProcessor(db)}
	return dbController, nil
}

func (db MySQLDB) Select(dataType interface{}) clauses.SelectClause {
	return &MySQLSelectClause{
		dbpool:     db.dbpool,
		Processor:  db.Processor,
		primaryKey: db.parser.ExtractPrimaryKey(dataType),
		tableName:  db.parser.GetTableName(dataType),
		columns:    db.parser.ExtractAllColumnsFromStructOrSlice(dataType, false),
		condition:  newCondition(db.parser),
	}
}

func (db MySQLDB) Insert(dataType interface{}) clauses.InsertClause {
	return &MySQLInsertClause{
		dbpool:    db.dbpool,
		Processor: db.Processor,
		tableName: db.parser.GetTableName(dataType),
		parser:    db.parser,
	}
}

func (db MySQLDB) Delete(dataType interface{}) clauses.DeleteClause {
	return &MySQLDeleteClause{
		dbpool:    db.dbpool,
		Processor: db.Processor,
		tableName: db.parser.GetTableName(dataType),
		parser:    db.parser,
		condition: newCondition(db.parser),
	}
}

func (db MySQLDB) Update(dataType interface{}) clauses.UpdateClause {
	return &MySQLUpdateClause{
		dbpool:    db.dbpool,
		Processor: db.Processor,
		tableName: db.parser.GetTableName(dataType),
		parser:    db.parser,
		condition: newCondition(db.parser),
	}
}
