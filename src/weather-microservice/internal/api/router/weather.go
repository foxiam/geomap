package router

import (
	"weather-microservice/internal/api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type WeatherService struct {
	app     *fiber.App
	handler *handler.WeatherHandler
}

func NewWeatherService(app *fiber.App) *WeatherService {
	return &WeatherService{app: app, handler: handler.NewWeatherHandler()}
}

func (s *WeatherService) Router() {

	api := s.app.Group("/api", logger.New())

	api.Get("/all", s.handler.GetAllCities)
	api.Get("/:name", s.handler.GetByName)
	api.Post("/filter", s.handler.GetByFilter)

}
