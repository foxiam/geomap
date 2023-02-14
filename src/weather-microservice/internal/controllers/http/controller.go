package v1

import "github.com/gofiber/fiber/v2"

func base(c *fiber.Ctx) error {
	return c.SendString("Hello, geomap 👋!")
}
