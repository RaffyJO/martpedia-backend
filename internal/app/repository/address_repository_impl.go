package repository

import (
	"martpedia-backend/internal/app/model/domain"

	"gorm.io/gorm"
)

type AddressRepositoryImpl struct {
	DB *gorm.DB
}

func NewAddressRepositoryImpl(db *gorm.DB) AddressRepository {
	return &AddressRepositoryImpl{
		DB: db,
	}
}

func (repository *AddressRepositoryImpl) Save(address domain.Address) (domain.Address, error) {
	address = domain.Address{
		Label:           address.Label,
		AddressLine1:    address.AddressLine1,
		AddressLine2:    address.AddressLine2,
		City:            address.City,
		State:           address.State,
		PostalCode:      address.PostalCode,
		Country:         address.Country,
		User:            address.User,
		AddressableID:   address.AddressableID,
		AddressableType: address.AddressableType,
	}

	response := repository.DB.Create(&address)
	if response != nil {
		return address, response.Error
	}

	return address, nil
}
