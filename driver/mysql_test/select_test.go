package mysql_test

import (
	"testing"

	"github.com/HarrisonEagle/HawkORM"
	"github.com/stretchr/testify/assert"
)

func TestMySQLSelectQuery(t *testing.T) {
	testDB, err := HawkORM.OpenMySQL("test")
	assert.NoError(t, err)
	query := testDB.Select(&User{}).WhereOr(&User{ID: "userId", Name: "testname"}).GetSQLQuery()
	assert.Equal(t, "SELECT id, name, email, created_at, updated_at FROM users WHERE (id = \"userId\" OR name = \"testname\")", query)
}
