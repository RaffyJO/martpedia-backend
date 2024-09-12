package repository

import "martpedia-backend/internal/app/model/domain"

type AddressRepository interface {
	Save(address domain.Address) (domain.Address, error)
	Update(address domain.Address) (domain.Address, error)
	Delete(address domain.Address) error
	FindById(id int) (domain.Address, error)
	FindAll(id int) ([]domain.Address, error)
}
