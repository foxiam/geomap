package internal

import (
	"context"
	"fmt"
	"log"

	"user-microservice/internal/api/router"
	"user-microservice/internal/config"
	"user-microservice/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func Run() error {
	config.InitEnvConfigs()

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		config.EnvConfig.DBUsername,
		config.EnvConfig.DBPassword,
		config.EnvConfig.DBHost,
		config.EnvConfig.DBPort,
		config.EnvConfig.DBName,
	)
	err := database.InitDB(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	userService := router.NewUserService(app)
	userService.Router()
	app.Listen(":3000")

	return nil
}
