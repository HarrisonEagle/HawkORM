package mysql_test

import (
	"testing"

	"github.com/HarrisonEagle/HawkORM"
	"github.com/stretchr/testify/assert"
)

func TestMySQLDelete(t *testing.T) {
	testDB, err := HawkORM.OpenMySQL("test")
	assert.NoError(t, err)
	// equal to: DELETE FROM users WHERE id = "testuserid4" AND email = "test@gmail.com"
	query := testDB.Delete(&User{}).Where(&User{ID: "testuserid4", Email: "test@gmail.com"}).GetSQLQuery()
	assert.Equal(t, "DELETE FROM users WHERE id = \"testuserid4\" AND email = \"test@gmail.com\"", query)
}
