package controller

import (
	"net/http"

	"loyalty_central/internal/models"
	"loyalty_central/internal/service"
	"loyalty_central/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	service service.AdminService
}

func NewAdminController(service service.AdminService) *AdminController {
	return &AdminController{
		service: service,
	}
}

func (c *AdminController) CreateAdmin(ctx *fiber.Ctx) error {
	var admin models.Admin

	err := ctx.BodyParser(&admin)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	createdAdmin, err := c.service.CreateAdmin(&admin)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.JSON(createdAdmin)
}

func (c *AdminController) GetAllAdmins(ctx *fiber.Ctx) error {
	admins, err := c.service.GetAllAdmins()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.JSON(admins)
}

func (c *AdminController) GetAdminByUsername(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	admin, err := c.service.GetAdminByUsername(username)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.JSON(admin)
}

func (c *AdminController) UpdateAdmin(ctx *fiber.Ctx) error {
	var admin models.Admin
	adminID, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	err = ctx.BodyParser(&admin)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	updatedAdmin, err := c.service.UpdateAdmin(&admin, uint(adminID))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.JSON(updatedAdmin)
}

func (c *AdminController) DeleteAdmin(ctx *fiber.Ctx) error {
	adminID, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	err = c.service.DeleteAdmin(uint(adminID))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.SendStatus(http.StatusOK)
}

func (c *AdminController) UndeleteAdminByID(ctx *fiber.Ctx) error {
	adminID, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	err = c.service.UndeleteAdminByID(uint(adminID))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.SendStatus(http.StatusOK)
}

func (c *AdminController) GetDeletedAdmins(ctx *fiber.Ctx) error {
	admins, err := c.service.GetDeletedAdmins()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.JSON(admins)
}
