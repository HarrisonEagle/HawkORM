package clauses

import "database/sql"

type InsertClause interface {
	SetData(data interface{}) InsertClause
	Exec() (sql.Result, error)
}
