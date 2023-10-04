package clauses

import "database/sql"

type InsertClause interface {
	SetData(data interface{}) InsertClause
	GetSQLQuery() string
	Exec() (sql.Result, error)
}
