package mysql_test

import "time"

type User struct {
	ID        string `hawkorm:"primarykey"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
