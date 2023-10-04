package clauses

type SelectClause interface {
	Limit(number int) SelectClause
	OrderBy(orderBy []string) SelectClause
	Where(condition interface{}) SelectClause
	WhereNot(condition interface{}) SelectClause
	WhereOr(condition interface{}) SelectClause
	GetSQLQuery() string
	All(target interface{}) error
	First(target interface{}) error
	Last(target interface{}) error
}
