package service

import (
	"errors"
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

func (service *AddressServiceImpl) Update(request web.UserAddressRequest, userId int) (web.UserAddressResponse, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return web.UserAddressResponse{}, err // Return the validation error
	}

	address, err := service.AddressRepository.FindById(request.ID)
	if err != nil {
		return web.UserAddressResponse{}, err // Return the error if the address is not found
	}

	if address.User.ID != userId {
		return web.UserAddressResponse{}, errors.New("Don't have permission to update this address")
	}

	address.Label = request.Label
	address.AddressLine1 = request.AddressLine1
	address.AddressLine2 = request.AddressLine2
	address.City = request.City
	address.State = request.State
	address.PostalCode = request.PostalCode
	address.Country = request.Country

	result, err := service.AddressRepository.Update(address)
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

func (service *AddressServiceImpl) Delete(id int, userId int) error {
	address, err := service.AddressRepository.FindById(id)
	if err != nil {
		return err // Return the error if the address is not found
	}

	if address.User.ID != userId {
		return errors.New("Don't have permission to delete this address") // Return the error if the user is not the owner of the address
	}

	err = service.AddressRepository.Delete(address)
	if err != nil {
		return err // Return the error if failed to delete the address
	}

	return nil
}

func (service *AddressServiceImpl) FindById(id int, userId int) (web.UserAddressResponse, error) {
	address, err := service.AddressRepository.FindById(id)
	if err != nil {
		return web.UserAddressResponse{}, err // Return the error if the address is not found
	}

	if address.AddressableID != userId {
		return web.UserAddressResponse{}, errors.New("Don't have permission to view this address")
	}

	return web.UserAddressResponse{
		ID:              address.ID,
		Label:           address.Label,
		AddressLine1:    address.AddressLine1,
		AddressLine2:    address.AddressLine2,
		City:            address.City,
		State:           address.State,
		PostalCode:      address.PostalCode,
		Country:         address.Country,
		AddressableID:   address.AddressableID,
		AddressableType: address.AddressableType,
	}, nil
}

func (service *AddressServiceImpl) FindAll(userId int) ([]web.UserAddressResponse, error) {
	addresses, err := service.AddressRepository.FindAll(userId)
	if err != nil {
		return []web.UserAddressResponse{}, err // Return the error if the address is not found
	}

	var response []web.UserAddressResponse
	for _, address := range addresses {
		response = append(response, web.UserAddressResponse{
			ID:              address.ID,
			Label:           address.Label,
			AddressLine1:    address.AddressLine1,
			AddressLine2:    address.AddressLine2,
			City:            address.City,
			State:           address.State,
			PostalCode:      address.PostalCode,
			Country:         address.Country,
			AddressableID:   address.AddressableID,
			AddressableType: address.AddressableType,
		})
	}

	return response, nil
}
