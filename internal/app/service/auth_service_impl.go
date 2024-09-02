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

func (service *AuthServiceImpl) Register(request web.UserRegisterRequest) (string, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return "", err // Return the validation error
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err // Return the error if hashing the password fails
	}

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
		Name:     request.Name,
		Phone:    request.Phone,
		Photo:    request.Photo,
	}

	result, err := service.AuthRepository.Save(user)
	if err != nil {
		return "", err // Return the error if the user already exists
	}

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

	return tokenString, nil
}

func (service *AuthServiceImpl) Login(request web.UserLoginRequest) (string, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return "", err // Return the validation error
	}

	user, err := service.AuthRepository.FindByEmailAndPassword(request.Email, request.Password)
	if err != nil {
		return "", err // Return the error if the user is not found or password is incorrect
	}

	// Create JWT token and include user data in the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"email":    user.Email,
		"name":     user.Name,
		"phone":    user.Phone,
		"photo":    user.Photo,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err // Return the error if signing the token fails
	}

	return tokenString, nil
}
