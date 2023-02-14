package app

import (
	"github.com/gofiber/fiber/v2"
	"weather-microservice/internal/controllers/http"
)

func Run() {
	/*
		l := logger.New(cfg.Log.Level)

			// Repository
			pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
			if err != nil {
				l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
			}
			defer pg.Close()
	*/

	handler := fiber.New()

	v1.NewRouter(handler)

	handler.Listen(":3000")

}
