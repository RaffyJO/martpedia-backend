package domain

import "time"

type Address struct {
	ID              int       `gorm:"column:id"`
	Label           string    `gorm:"column:label"`
	AddressLine1    string    `gorm:"column:address_line_1"`
	AddressLine2    string    `gorm:"column:address_line_2"`
	City            string    `gorm:"column:city"`
	State           string    `gorm:"column:state"`
	PostalCode      string    `gorm:"column:postal_code"`
	Country         string    `gorm:"column:country"`
	AddressableID   int       `gorm:"column:addressable_id"`
	AddressableType string    `gorm:"column:addressable_type"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	User            User      `gorm:"foreignKey:addressable_id;references:id"`
}
