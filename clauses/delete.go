package clauses

import "database/sql"

type DeleteClause interface {
	Where(condition interface{}) DeleteClause
	WhereOr(condition interface{}) DeleteClause
	WhereNot(condition interface{}) DeleteClause
	Limit(number int) DeleteClause
	Exec() (sql.Result, error)
}
