package driver

import "github.com/HarrisonEagle/HawkORM/clauses"

type Database interface {
	Select(dataType interface{}) clauses.SelectClause
	Insert(dataType interface{}) clauses.InsertClause
	Delete(dataType interface{}) clauses.DeleteClause
	Update(dataType interface{}) clauses.UpdateClause
}
