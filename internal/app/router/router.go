package router

import (
	"martpedia-backend/internal/app/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(authController controller.AuthController) *fiber.App {
	router := fiber.New()

	router.Post("/api/register", authController.Register)
	router.Post("/api/login", authController.Login)

	return router
}
