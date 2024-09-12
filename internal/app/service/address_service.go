package service

import "martpedia-backend/internal/app/model/web"

type AddressService interface {
	Create(request web.UserAddressRequest) (web.UserAddressResponse, error)
}
