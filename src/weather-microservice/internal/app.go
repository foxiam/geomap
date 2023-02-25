package internal

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"log"
	"weather-microservice/internal/api/router"
	configs "weather-microservice/pkg"
	"weather-microservice/pkg/database"
)

func Run() {

	configs.InitEnvConfigs()

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		configs.EnvConfigs.DBUsername,
		configs.EnvConfigs.DBPassword,
		configs.EnvConfigs.DBHost,
		configs.EnvConfigs.DBPort,
		configs.EnvConfigs.DBName,
	)
	err := database.InitDB(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	WeatherService := router.NewWeatherService(app)
	WeatherService.Router()
	app.Listen(configs.EnvConfigs.LocalServerPort)

}
