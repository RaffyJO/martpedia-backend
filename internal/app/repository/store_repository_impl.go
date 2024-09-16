package repository

import (
	"martpedia-backend/internal/app/model/domain"

	"gorm.io/gorm"
)

type StoreRepositoryImpl struct {
	DB *gorm.DB
}

func NewStoreRepositoryImpl(db *gorm.DB) StoreRepository {
	return &StoreRepositoryImpl{
		DB: db,
	}
}

func (repository *StoreRepositoryImpl) Save(store domain.Store) (domain.Store, error) {
	store = domain.Store{
		Name:        store.Name,
		OwnerID:     store.OwnerID,
		AddressID:   store.AddressID,
		Description: store.Description,
		Photo:       store.Photo,
	}

	response := repository.DB.Create(&store)
	if response != nil {
		return store, response.Error
	}

	return store, nil
}
