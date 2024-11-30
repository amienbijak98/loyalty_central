package controller

import (
	"net/http"

	"loyalty_central/internal/dto"
	"loyalty_central/internal/models"
	"loyalty_central/internal/service"
	"loyalty_central/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	service service.CustomerService
}

func NewCustomerController(service service.CustomerService) *CustomerController {
	return &CustomerController{
		service: service,
	}
}

func (c *CustomerController) CreateCustomer(ctx *fiber.Ctx) error {
	var customer models.Customer

	err := ctx.BodyParser(&customer)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	result, err := c.service.CreateCustomer(&customer)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.JSON(result)
}

func (c *CustomerController) GetCustomerByPhoneNumber(ctx *fiber.Ctx) error {
	phoneNumber := ctx.Params("phone_number")

	result, err := c.service.GetCustomerByPhoneNumber(phoneNumber)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.JSON(result)
}

func (c *CustomerController) UpdateCustomer(ctx *fiber.Ctx) error {
	var customer models.Customer

	customerID, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	err = ctx.BodyParser(&customer)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	result, err := c.service.UpdateCustomer(&customer, uint(customerID))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.JSON(result)
}

func (c *CustomerController) DeleteCustomer(ctx *fiber.Ctx) error {
	customerID, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	err = c.service.DeleteCustomer(uint(customerID))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.SendStatus(http.StatusOK)
}

func (c *CustomerController) GetAllCustomers(ctx *fiber.Ctx) error {
	result, err := c.service.GetAllCustomers()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.JSON(result)
}

func (c *CustomerController) GetCustomerByID(ctx *fiber.Ctx) error {
	customerID, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	customer, err := c.service.GetCustomerByID(uint(customerID))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	// Mapping data dari models.Customer ke dto.CustomerDTO
	customerDTO := dto.CustomerDTO{
		ID:          customer.ID,
		Name:        customer.Name,
		PhoneNumber: customer.PhoneNumber,
		TotalPoints: customer.TotalPoints,
	}

	return ctx.JSON(customerDTO)
}

func (c *CustomerController) UndeleteCustomerByID(ctx *fiber.Ctx) error {
	customerID, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewErrorResponse(err))
	}

	err = c.service.UndeleteCustomerByID(uint(customerID))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.SendStatus(http.StatusOK)
}

func (c *CustomerController) GetDeletedCustomers(ctx *fiber.Ctx) error {
	result, err := c.service.GetDeletedCustomers()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewErrorResponse(err))
	}

	return ctx.JSON(result)
}
