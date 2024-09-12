package controller

import (
	"fmt"
	"martpedia-backend/internal/app/model/domain"
	"martpedia-backend/internal/app/model/web"
	"martpedia-backend/internal/app/service"
	"strconv"

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

func (controller *AddressControllerImpl) Update(ctx *fiber.Ctx) error {
	userAddressRequest := web.UserAddressRequest{}
	if err := ctx.BodyParser(&userAddressRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	addressId := ctx.Params("id")
	id, err := strconv.Atoi(addressId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	userAddressRequest.ID = int(id)
	user := ctx.Locals("user").(domain.User)

	response, err := controller.AddressService.Update(userAddressRequest, user.ID)
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
		"message": "Address updated successfully",
		"data":    response,
	})
}

func (controller *AddressControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	user := ctx.Locals("user").(domain.User)

	err = controller.AddressService.Delete(id, user.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Address deleted successfully",
	})
}

func (controller *AddressControllerImpl) FindById(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	user := ctx.Locals("user").(domain.User)

	response, err := controller.AddressService.FindById(int(id), user.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Address found successfully",
		"data":    response,
	})
}

func (controller *AddressControllerImpl) FindAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(domain.User)

	response, err := controller.AddressService.FindAll(user.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Addresses found successfully",
		"data":    response,
	})
}
