package controller

import (
	"fmt"
	"martpedia-backend/internal/app/model/web"
	"martpedia-backend/internal/app/service"

	"github.com/go-playground/validator/v10"
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

	response, err := controller.AuthService.Register(userRegisterRequest)
	if err != nil {
		// Jika error adalah error validasi dari validator
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			var errors []string
			for _, err := range validationErrs {
				// Menyusun pesan error spesifik untuk setiap field
				errors = append(errors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()))
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": errors,
			})
		}

		// Jika error lain
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User registered successfully",
		"data": fiber.Map{
			"token": response,
		},
	})
}

func (controller *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	userLoginRequest := web.UserLoginRequest{}
	if err := ctx.BodyParser(&userLoginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	response, err := controller.AuthService.Login(userLoginRequest)
	if err != nil {
		// Jika error adalah error validasi dari validator
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			var errors []string
			for _, err := range validationErrs {
				// Menyusun pesan error spesifik untuk setiap field
				errors = append(errors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()))
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": errors,
			})
		}

		// Jika error lain
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User logged in successfully",
		"data": fiber.Map{
			"token": response,
		},
	})
}
