package router

import (
	"user-microservice/internal/api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type userServer struct {
	app     *fiber.App
	handler *handler.UserHandler
}

func NewUserServer(app *fiber.App, handler *handler.UserHandler) *userServer {
	return &userServer{app: app, handler: handler}
}

func (s *userServer) Router() {

	api := s.app.Group("/api", logger.New())
	//Auth
	auth := api.Group("/auth")
	auth.Post("/login", s.handler.SignIn)
	auth.Post("/registration", s.handler.SingUp)

	// User
	user := api.Group("/user")
	user.Get("/all", s.handler.GetAllUsers)
	user.Get("/:id", s.handler.GetUser)
	//user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	//user.Delete("/:id", middleware.Protected(), s.handler.DeleteUser)
}
