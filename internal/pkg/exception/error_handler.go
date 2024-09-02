package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if notFoundError(ctx, err) {
		return nil
	}

	if validationErrors(ctx, err) {
		return nil
	}

	internalServerError(ctx, err)
	return nil
}

func validationErrors(ctx *fiber.Ctx, err error) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(fiber.Map{
			"status":  "BAD REQUEST",
			"message": exception.Error(),
		})
		return true
	}
	return false
}

func notFoundError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		ctx.Status(fiber.StatusNotFound)
		ctx.JSON(fiber.Map{
			"status":  "NOT FOUND",
			"message": exception.Error,
		})
		return true
	}
	return false
}

func internalServerError(ctx *fiber.Ctx, err error) {
	exception, ok := err.(error)
	if ok {
		ctx.Status(fiber.StatusInternalServerError)
		ctx.JSON(fiber.Map{
			"status":  "INTERNAL SERVER ERROR",
			"message": exception.Error(),
		})
		return
	}
}
