package clauses

import "database/sql"

type UpdateClause interface {
	SetData(data interface{}) UpdateClause
	Where(condition interface{}) UpdateClause
	WhereOr(condition interface{}) UpdateClause
	WhereNot(condition interface{}) UpdateClause
	Exec() (sql.Result, error)
	OrderBy(orderBy []string) UpdateClause
	Limit(number int) UpdateClause
}
