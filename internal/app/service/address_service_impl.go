package service

import (
	"martpedia-backend/internal/app/model/domain"
	"martpedia-backend/internal/app/model/web"
	"martpedia-backend/internal/app/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AddressServiceImpl struct {
	AddressRepository repository.AddressRepository
	DB                *gorm.DB
	validate          *validator.Validate
}

func NewAddressServiceImpl(addressRepository repository.AddressRepository, db *gorm.DB, validate *validator.Validate) AddressService {
	return &AddressServiceImpl{
		AddressRepository: addressRepository,
		DB:                db,
		validate:          validate,
	}
}

func (service *AddressServiceImpl) Create(request web.UserAddressRequest) (web.UserAddressResponse, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return web.UserAddressResponse{}, err // Return the validation error
	}

	address := domain.Address{
		Label:           request.Label,
		AddressLine1:    request.AddressLine1,
		AddressLine2:    request.AddressLine2,
		City:            request.City,
		State:           request.State,
		PostalCode:      request.PostalCode,
		Country:         request.Country,
		AddressableID:   request.ID,
		AddressableType: request.AddressableType,
	}

	result, err := service.AddressRepository.Save(address)
	if err != nil {
		return web.UserAddressResponse{}, err // Return the error if the address already exists
	}

	return web.UserAddressResponse{
		ID:              result.ID,
		Label:           result.Label,
		AddressLine1:    result.AddressLine1,
		AddressLine2:    result.AddressLine2,
		City:            result.City,
		State:           result.State,
		PostalCode:      result.PostalCode,
		Country:         result.Country,
		AddressableID:   result.AddressableID,
		AddressableType: result.AddressableType,
	}, nil
}
