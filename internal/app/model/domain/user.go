package domain

import "time"

type User struct {
	ID        int       `gorm:"column:id"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Phone     string    `gorm:"column:phone"`
	Photo     string    `gorm:"column:photo"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Addresses []Address `gorm:"foreignKey:addressable_id;references:id"`
}
