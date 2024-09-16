package domain

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name        string  `json:"name"`
	OwnerID     int     `json:"owner_id"`
	AddressID   int     `json:"address_id"`
	Description string  `json:"description"`
	Photo       string  `json:"photo"`
	User        User    `gorm:"foreignKey:owner_id;references:id"`
	Address     Address `gorm:"foreignKey:address_id;references:id"`
}
