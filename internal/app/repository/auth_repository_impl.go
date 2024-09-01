package repository

import (
	"martpedia-backend/internal/app/model/domain"
	"martpedia-backend/internal/pkg/helper"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

func (repository *AuthRepositoryImpl) Save(user domain.User) domain.User {
	user = domain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		Phone:    user.Phone,
	}

	response := repository.DB.Create(&user)
	helper.PanicIfError(response.Error)

	return user
}
