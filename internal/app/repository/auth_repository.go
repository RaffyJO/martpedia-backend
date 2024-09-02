package repository

import (
	"martpedia-backend/internal/app/model/domain"
)

type AuthRepository interface {
	Save(user domain.User) (domain.User, error)
	FindByEmailAndPassword(email string, password string) (domain.User, error)
}
