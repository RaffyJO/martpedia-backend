package controller

import "github.com/gofiber/fiber/v2"

type StoreController interface {
	Create(ctx *fiber.Ctx) error
}
