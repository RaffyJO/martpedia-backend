// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"martpedia-backend/internal/app/controller"
	"martpedia-backend/internal/app/db"
	"martpedia-backend/internal/app/repository"
	"martpedia-backend/internal/app/router"
	"martpedia-backend/internal/app/service"
)

// Injectors from injector.go:

// InitializedServer creates a new Fiber application with dependencies injected
func InitializedServer() *fiber.App {
	gormDB := db.NewDB()
	authRepository := repository.NewAuthRepositoryImpl(gormDB)
	validate := ProvideValidator()
	authService := service.NewAuthServiceImpl(authRepository, gormDB, validate)
	authController := controller.NewAuthControllerImpl(authService)
	addressRepository := repository.NewAddressRepositoryImpl(gormDB)
	addressService := service.NewAddressServiceImpl(addressRepository, gormDB, validate)
	addressController := controller.NewAddressControllerImpl(addressService)
	storeRepository := repository.NewStoreRepositoryImpl(gormDB)
	storeService := service.NewStoreServiceImpl(storeRepository, gormDB, validate, addressService)
	storeController := controller.NewStoreControllerImpl(storeService)
	app := router.NewRouter(authController, addressController, storeController)
	return app
}

// injector.go:

var authSet = wire.NewSet(repository.NewAuthRepositoryImpl, service.NewAuthServiceImpl, controller.NewAuthControllerImpl)

var addressSet = wire.NewSet(repository.NewAddressRepositoryImpl, service.NewAddressServiceImpl, controller.NewAddressControllerImpl)

var storeSet = wire.NewSet(repository.NewStoreRepositoryImpl, service.NewStoreServiceImpl, controller.NewStoreControllerImpl)

// ProvideValidator creates and returns a new validator instance
func ProvideValidator() *validator.Validate {
	v := validator.New()
	return v
}
