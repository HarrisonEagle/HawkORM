package mysql_test

import (
	"testing"

	"github.com/HarrisonEagle/HawkORM"
	"github.com/stretchr/testify/assert"
)

func TestMySQLInsert(t *testing.T) {
	testDB, err := HawkORM.OpenMySQL("test")
	assert.NoError(t, err)
	insertTest := User{ID: "userid1", Name: "user1", Email: "test@gmail.com"}
	query := testDB.Insert(&User{}).SetData(insertTest).GetSQLQuery()
	assert.Equal(t, "INSERT INTO users (id, name, email) VALUES (\"userid1\", \"user1\", \"test@gmail.com\")", query)
}

func TestMySQLBulkInsert(t *testing.T) {
	testDB, err := HawkORM.OpenMySQL("test")
	assert.NoError(t, err)
	insertTest := []User{
		{ID: "userid1", Name: "user1", Email: "test@gmail.com"},
		{ID: "userid2", Name: "user2", Email: "test2@gmail.com"},
	}
	// equal to: INSERT INTO users (id, name, email) VALUES ("userid1", "user1", "test@gmail.com"), ("userid2", "user2", "test2@gmail.com")
	query := testDB.Insert(&User{}).SetData(insertTest).GetSQLQuery()
	assert.Equal(t, "INSERT INTO users (id, name, email) VALUES (\"userid1\", \"user1\", \"test@gmail.com\"), (\"userid2\", \"user2\", \"test2@gmail.com\")", query)
}
