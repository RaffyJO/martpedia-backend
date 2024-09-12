package controller

import (
	"fmt"
	"martpedia-backend/internal/app/model/domain"
	"martpedia-backend/internal/app/model/web"
	"martpedia-backend/internal/app/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AddressControllerImpl struct {
	AddressService service.AddressService
}

func NewAddressControllerImpl(addressService service.AddressService) AddressController {
	return &AddressControllerImpl{
		AddressService: addressService,
	}
}

func (controller *AddressControllerImpl) Create(ctx *fiber.Ctx) error {
	userAddressRequest := web.UserAddressRequest{}
	if err := ctx.BodyParser(&userAddressRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	// Get the user information from the context (which is stored by the middleware `RequiredAuth`)
	user := ctx.Locals("user").(domain.User)

	// Include the user ID in the request body
	userAddressRequest.ID = user.ID

	response, err := controller.AddressService.Create(userAddressRequest)
	if err != nil {
		// If the error is a validation error from the validator
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			var errors []string
			for _, err := range validationErrs {
				// Display specific error message for each field
				errors = append(errors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()))
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": errors,
			})
		}

		// If an another error occurs
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Address created successfully",
		"data":    response,
	})
}
