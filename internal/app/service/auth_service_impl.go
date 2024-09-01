package service

import (
	"martpedia-backend/internal/app/model/domain"
	"martpedia-backend/internal/app/model/web"
	"martpedia-backend/internal/app/repository"
	"martpedia-backend/internal/pkg/helper"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *gorm.DB
	validate       *validator.Validate
}

func NewAuthServiceImpl(authRepository repository.AuthRepository, db *gorm.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             db,
		validate:       validate,
	}
}

func (service *AuthServiceImpl) Register(request web.UserRegisterRequest) string {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
		Name:     request.Name,
		Phone:    request.Phone,
		Photo:    request.Photo,
	}

	result := service.AuthRepository.Save(user)

	// Create JWT token and include user data in the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      result.ID,
		"username": result.Username,
		"email":    result.Email,
		"name":     result.Name,
		"phone":    result.Phone,
		"photo":    result.Photo,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	helper.PanicIfError(err)

	return tokenString
}
