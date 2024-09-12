package service

import "martpedia-backend/internal/app/model/web"

type AddressService interface {
	Create(request web.UserAddressRequest) (web.UserAddressResponse, error)
	Update(request web.UserAddressRequest, userId int) (web.UserAddressResponse, error)
	Delete(id int, userId int) error
	FindById(id int, userId int) (web.UserAddressResponse, error)
	FindAll(userId int) ([]web.UserAddressResponse, error)
}
