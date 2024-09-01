package controller

import (
	"martpedia-backend/internal/app/model/web"
	"martpedia-backend/internal/app/service"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthControllerImpl(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	userRegisterRequest := web.UserRegisterRequest{}
	if err := ctx.BodyParser(&userRegisterRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	response := controller.AuthService.Register(userRegisterRequest)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User registered successfully",
		"data": fiber.Map{
			"token": response,
		},
	})
}
