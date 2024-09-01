package domain

import "time"

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Name      string
	Phone     string
	Photo     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
