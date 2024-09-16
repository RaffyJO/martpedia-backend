package service

import (
	"martpedia-backend/internal/app/model/domain"
	"martpedia-backend/internal/app/model/web"
	"martpedia-backend/internal/app/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type StoreServiceImpl struct {
	StoreRepository repository.StoreRepository
	DB              *gorm.DB
	validate        *validator.Validate
	AddressService  AddressService
}

func NewStoreServiceImpl(storeRepository repository.StoreRepository, db *gorm.DB, validate *validator.Validate, addressService AddressService) StoreService {
	return &StoreServiceImpl{
		StoreRepository: storeRepository,
		DB:              db,
		validate:        validate,
		AddressService:  addressService,
	}
}

func (service *StoreServiceImpl) Create(request web.StoreCreateRequest) (web.StoreCreateResponse, error) {
	err := service.validate.Struct(request.StoreRequest)
	if err != nil {
		return web.StoreCreateResponse{}, err // Return the validation error if the store request is invalid
	}

	err = service.validate.Struct(request.AddressRequest)
	if err != nil {
		return web.StoreCreateResponse{}, err // Return the validation error if the address request is invalid
	}

	addressRequest := web.UserAddressRequest{
		Label:           request.AddressRequest.Label,
		AddressLine1:    request.AddressRequest.AddressLine1,
		AddressLine2:    request.AddressRequest.AddressLine2,
		City:            request.AddressRequest.City,
		State:           request.AddressRequest.State,
		PostalCode:      request.AddressRequest.PostalCode,
		Country:         request.AddressRequest.Country,
		AddressableID:   request.AddressRequest.AddressableID,
		AddressableType: "store",
	}

	address, err := service.AddressService.Create(addressRequest)
	if err != nil {
		return web.StoreCreateResponse{}, err // Return the error if the address creation fails
	}

	store := domain.Store{
		Name:        request.StoreRequest.Name,
		OwnerID:     request.StoreRequest.OwnerID,
		AddressID:   address.ID,
		Description: request.StoreRequest.Description,
		Photo:       request.StoreRequest.Photo,
	}

	result, err := service.StoreRepository.Save(store)
	if err != nil {
		return web.StoreCreateResponse{}, err // Return the error if the store creation fails
	}

	return web.StoreCreateResponse{
		StoreResponse: web.StoreResponse{
			ID:          int(result.ID),
			Name:        result.Name,
			OwnerID:     result.OwnerID,
			AddressID:   result.AddressID,
			Description: result.Description,
			Photo:       result.Photo,
			CreatedAt:   result.CreatedAt.Format("2006-01-02 15:04:05"),
		},
		AddressResponse: web.UserAddressResponse{
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
		},
	}, nil
}
