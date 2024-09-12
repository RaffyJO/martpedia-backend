package repository

import "martpedia-backend/internal/app/model/domain"

type AddressRepository interface {
	Save(address domain.Address) (domain.Address, error)
}
