package internal

import (
	"context"
	"log"

	"user-microservice/internal/api/handler"
	"user-microservice/internal/api/router"
	"user-microservice/internal/config"
	"user-microservice/internal/repository"
	"user-microservice/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func Run() error {
	config.InitEnvConfigs()

	pool, err := database.NewPostgresDB(
		context.Background(),
		database.Config{
			Host:     config.EnvConfig.DBHost,
			Port:     config.EnvConfig.DBPort,
			Username: config.EnvConfig.DBUsername,
			Name:     config.EnvConfig.DBName,
			Password: config.EnvConfig.DBPassword,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	app := fiber.New()

	userRepo := repository.NewUserRepository(pool)
	userHandler := handler.NewUserHandler(userRepo)
	userServer := router.NewUserServer(app, userHandler)

	userServer.Router()
	app.Listen(":3000")

	return nil
}
