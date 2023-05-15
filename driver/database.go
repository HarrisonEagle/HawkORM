package driver

import "github.com/HarrisonOwl/GABAORM/clauses"

type Database interface {
	Select(dataType interface{}) clauses.SelectClause
}
