package repository

import (
	"martpedia-backend/internal/app/model/domain"

	"golang.org/x/crypto/bcrypt"
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

func (repository *AuthRepositoryImpl) Save(user domain.User) (domain.User, error) {
	user = domain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		Phone:    user.Phone,
	}

	response := repository.DB.Create(&user)
	if response != nil {
		return user, response.Error
	}

	return user, nil
}

func (repository *AuthRepositoryImpl) FindByEmailAndPassword(email string, password string) (domain.User, error) {
	var user domain.User

	response := repository.DB.Where("email = ?", email).First(&user)
	if response.Error != nil {
		return user, response.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}
