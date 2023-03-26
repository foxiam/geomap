package handler

import (
	"context"
	"fmt"
	"weather-microservice/internal/model"
	"weather-microservice/internal/repository"
	"weather-microservice/pkg/database"

	"github.com/gofiber/fiber/v2"
)

type WeatherHandler struct {
	WeatherRepository *repository.WeatherRepository
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{WeatherRepository: repository.NewWeatherRepository(database.GetPool())}
}

func (h *WeatherHandler) GetAllCities(c *fiber.Ctx) error {
	cities, err := h.WeatherRepository.GetAll(context.Background())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"data": cities})
}

func (h *WeatherHandler) GetByName(c *fiber.Ctx) error {
	name := c.Params("name")
	city, err := h.WeatherRepository.GetByName(context.Background(), name)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"data": city})
}

func (h *WeatherHandler) GetByFilter(c *fiber.Ctx) error {
	filter := new(model.WeatherFilter)
	if err := c.BodyParser(filter); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})

	}
	fmt.Println(filter)
	cities, err := h.WeatherRepository.FindAllByFilter(c.Context(), filter)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"data": cities})
}
