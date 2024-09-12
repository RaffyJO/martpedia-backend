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

func (repository *AddressRepositoryImpl) Update(address domain.Address) (domain.Address, error) {
	response := repository.DB.Model(&address).Where("id = ?", address.ID).Updates(address)
	if response != nil {
		return address, response.Error
	}

	return address, nil
}

func (repository *AddressRepositoryImpl) Delete(address domain.Address) error {
	response := repository.DB.Delete(&address)
	if response != nil {
		return response.Error
	}

	return nil
}

func (repository *AddressRepositoryImpl) FindById(id int) (domain.Address, error) {
	var address domain.Address
	response := repository.DB.First(&address, id)
	if response.Error != nil {
		return address, response.Error
	}

	return address, nil
}

func (repository *AddressRepositoryImpl) FindAll(id int) ([]domain.Address, error) {
	var addresses []domain.Address
	response := repository.DB.Where("addressable_id = ?", id).Order("id asc").Find(&addresses)
	if response.Error != nil {
		return addresses, response.Error
	}

	return addresses, nil
}
