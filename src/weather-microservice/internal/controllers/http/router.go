package v1

import "github.com/gofiber/fiber/v2"

func NewRouter(handler *fiber.App) {
	handler.Get("/", base)
}
