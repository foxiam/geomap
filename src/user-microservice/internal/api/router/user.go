package router

import (
	"user-microservice/internal/api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type userService struct {
	app     *fiber.App
	handler *handler.UserHandler
}

func NewUserService(app *fiber.App) *userService {
	return &userService{app: app, handler: handler.NewUserHandler()}
}

func (s *userService) Router() {

	api := s.app.Group("/api", logger.New())
	//Auth
	auth := api.Group("/auth")
	auth.Post("/login", s.handler.Login)

	// User
	user := api.Group("/user")
	user.Get("/all", s.handler.GetAllUsers)
	user.Get("/:id", s.handler.GetUser)
	user.Post("/", s.handler.CreateUser)
	//user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	//user.Delete("/:id", middleware.Protected(), handler.DeleteUser)
}
