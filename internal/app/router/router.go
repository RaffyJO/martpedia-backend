package router

import (
	"martpedia-backend/internal/app/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(authController controller.AuthController) *fiber.App {
	app := fiber.New()

	app.Post("/api/register", authController.Register)

	return app
}
