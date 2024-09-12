package controller

import "github.com/gofiber/fiber/v2"

type AddressController interface {
	Create(ctx *fiber.Ctx) error
}
