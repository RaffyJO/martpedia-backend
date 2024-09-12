//go:build wireinject
// +build wireinject

package main

import (
	"martpedia-backend/internal/app/controller"
	"martpedia-backend/internal/app/db"
	"martpedia-backend/internal/app/repository"
	"martpedia-backend/internal/app/router"
	"martpedia-backend/internal/app/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var authSet = wire.NewSet(
	repository.NewAuthRepositoryImpl,
	service.NewAuthServiceImpl,
	controller.NewAuthControllerImpl,
)

var addressSet = wire.NewSet(
	repository.NewAddressRepositoryImpl,
	service.NewAddressServiceImpl,
	controller.NewAddressControllerImpl,
)

// ProvideValidator creates and returns a new validator instance
func ProvideValidator() *validator.Validate {
	v := validator.New()
	return v
}

// InitializedServer creates a new Fiber application with dependencies injected
func InitializedServer() *fiber.App {
	wire.Build(
		db.NewDB,
		ProvideValidator,
		authSet,
		addressSet,
		router.NewRouter,
	)
	return nil
}
