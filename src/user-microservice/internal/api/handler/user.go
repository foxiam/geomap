package handler

import (
	"context"
	"strconv"

	"user-microservice/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler struct {
	userRepository *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepository: userRepo}
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func (h *UserHandler) validUser(id string, p string) bool {
	user, err := h.userRepository.FindByID(context.Background(), id)
	if err != nil {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userRepository.FindByID(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userRepository.FindAll(context.Background())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err, "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Users found", "data": users})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	type PasswordInput struct {
		Password string `json:"password"`
	}
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	if !h.validUser(id, pi.Password) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}
