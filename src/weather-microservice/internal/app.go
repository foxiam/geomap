package internal

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// Default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	WeatherService := router.NewWeatherService(app)
	WeatherService.Router()
	app.Listen(configs.EnvConfigs.LocalServerPort)

}
