package controller

import (
	"net/http"

	"loyalty_central/internal/dto"
	"loyalty_central/internal/service"
	"loyalty_central/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// AuthController provides authentication controller
type AuthController struct {
	service service.AuthService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

// Login authenticates a user and returns a JWT token if successful
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var login dto.LoginDTO

	err := ctx.BodyParser(&login)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	token, err := c.service.Login(login.Username, login.Password)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.JSON(fiber.Map{"token": token})
}
