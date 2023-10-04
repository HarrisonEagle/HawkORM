# HawkORM
A toy ORM for golang.
Currently only MySQL is supported.

## Installation
```shell
go get github.com/HarrisonEagle/HawkORM
```

## Progress
- **Select**
  - ✅select all objects
  - ✅select objects with conditions
  - ✅select first
  - ✅select last
  - 🔄left join, inner join, outer join
- **Insert**
  - ✅insert an object
  - ✅batch insert
  - 🔄insert with associations
- **Update**
  - ✅update with conditions
  - ✅batch update
- **Delete**
  - ✅delete with conditions
  - ✅delete an object
  - ✅delete with multiple objects
- **Transaction**
- 🔄Comming soon...

## Usage

For example, user data object defined like this:
```go
type User struct {
 ID        string `hawkorm:"primarykey"`
 Name      string
 Email     string
 CreatedAt time.Time
 UpdatedAt time.Time
}
```

### SELECT
```go
// connect to db
// remember to add parseTime=true to parse time correctlt
testDB, err := HawkORM.OpenMySQL("user:pass@localhost:3306/databasename?parseTime=true")
if err != nil{
	log.Fatalf("connect to db failed!")
}

// equal to: SELECT id, name, email, created_at, updated_at FROM users WHERE (id = "userId" OR name = "testname") ORDER BY id ASC LIMIT 1
var users []User
testDB.Select(&User{}).WhereOr(&User{ID: "userId", Name: "testname"}).First(&users)
```

### INSERT
```go
insertTest := []User{
 {ID: "userid1", Name: "user1", Email: "test@gmail.com"},
 {ID: "userid2", Name: "user2", Email: "test2@gmail.com"},
}
// equal to: INSERT INTO users (id, name, email) VALUES ("userid1", "user1", "test@gmail.com"), ("userid2", "user2", "test2@gmail.com")
res, err := testDB.Insert(&User{}).SetData(insertTest).Exec()
```

### UPDATE
```go
// equal to: UPDATE users SET name = "rename" WHERE id = "testuserid3" AND email = "test@gmail.com" 
_, err := testDB.Update(&User{}).Where(&User{ID: "testuserid3", Email: "test@gmail.com"}).SetData(User{Name: "rename"}).Exec()
```

### DELETE
```go
// equal to: DELETE FROM users WHERE id = "testuserid4" AND email = "test@gmail.com"
_, err := testDB.Delete(&User{}).Where(&User{ID: "testuserid4", Email: "test@gmail.com"}).Exec()
```
