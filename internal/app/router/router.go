package router

import (
	"martpedia-backend/internal/app/controller"
	"martpedia-backend/internal/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(authController controller.AuthController, addressController controller.AddressController, storeController controller.StoreController) *fiber.App {
	router := fiber.New()

	router.Use(middleware.DatabaseMiddleware())

	router.Post("/api/register", authController.Register)
	router.Post("/api/login", authController.Login)

	router.Post("/api/address", middleware.RequiredAuth, addressController.Create)
	router.Put("/api/address/:id", middleware.RequiredAuth, addressController.Update)
	router.Delete("/api/address/:id", middleware.RequiredAuth, addressController.Delete)
	router.Get("/api/address/:id", middleware.RequiredAuth, addressController.FindById)
	router.Get("/api/address", middleware.RequiredAuth, addressController.FindAll)

	router.Post("/api/store", middleware.RequiredAuth, storeController.Create)

	return router
}
