package GABAORM

import "time"

type Model struct {
	ID        uint		`gabaorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
