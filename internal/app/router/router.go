package router

import (
	"martpedia-backend/internal/app/controller"
	"martpedia-backend/internal/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(authController controller.AuthController, addressController controller.AddressController) *fiber.App {
	router := fiber.New()

	router.Use(middleware.DatabaseMiddleware())

	router.Post("/api/register", authController.Register)
	router.Post("/api/login", authController.Login)

	router.Post("/api/address", middleware.RequiredAuth, addressController.Create)

	return router
}
