package repository

import "martpedia-backend/internal/app/model/domain"

type StoreRepository interface {
	Save(store domain.Store) (domain.Store, error)
}
