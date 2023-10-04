package mysql_test

import (
	"testing"

	"github.com/HarrisonEagle/HawkORM"
	"github.com/stretchr/testify/assert"
)

func TestMySQLUpdate(t *testing.T) {
	testDB, err := HawkORM.OpenMySQL("test")
	assert.NoError(t, err)
	query := testDB.Update(&User{}).Where(&User{ID: "testuserid3", Email: "test@gmail.com"}).SetData(User{Name: "rename"}).GetSQLQuery()
	assert.Equal(t, "UPDATE users SET name = \"rename\" WHERE id = \"testuserid3\" AND email = \"test@gmail.com\"", query)
}
