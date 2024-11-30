package router

import (
	"loyalty_central/internal/controller"
	"loyalty_central/internal/middleware"
	"loyalty_central/internal/repository"
	"loyalty_central/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routing(app *fiber.App, db *gorm.DB) {

	api := app.Group("/api/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello this is Loyalty Central App API")
	})

	customerRoutes(api, db)

	adminRoutes(api, db)

	authRoutes(api, db)
}

func customerRoutes(api fiber.Router, db *gorm.DB) {
	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)

	api = api.Group("/customer")

	// Middleware untuk route yang diakses hanya oleh role "manager" dan "admin"
	allowedRoles := []string{"manager", "admin"}
	roleMiddleware := middleware.RoleMiddleware(allowedRoles)

	api.Put("/:id/undelete", roleMiddleware, customerController.UndeleteCustomerByID)
	api.Put("/:id", roleMiddleware, customerController.UpdateCustomer)
	api.Delete("/:id", roleMiddleware, customerController.DeleteCustomer)
	api.Get("/deleted", roleMiddleware, customerController.GetDeletedCustomers)
	api.Get("/:id", roleMiddleware, customerController.GetCustomerByID)

	// Middleware untuk route yang diakses hanya oleh role "manager", "admin", dan "casheer"
	allowedRoles = []string{"manager", "admin", "casheer"}
	roleMiddleware = middleware.RoleMiddleware(allowedRoles)
	api.Get("", roleMiddleware, customerController.GetAllCustomers)
	api.Post("/", roleMiddleware, customerController.CreateCustomer)
	api.Get("/phone/:phone_number", roleMiddleware, customerController.GetCustomerByPhoneNumber)
}

func adminRoutes(api fiber.Router, db *gorm.DB) {
	adminRepository := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepository)
	adminController := controller.NewAdminController(adminService)

	api.Post("/admin/register", adminController.CreateAdmin)
	api.Get("/admin/username/:username", adminController.GetAdminByUsername)
	api.Put("/admin/:id/undelete", adminController.UndeleteAdminByID)
	api.Put("/admin/:id", adminController.UpdateAdmin)
	api.Delete("/admin/:id", adminController.DeleteAdmin)
	api.Get("/admin", adminController.GetAllAdmins)
	api.Get("/admin/deleted", adminController.GetDeletedAdmins)
}

func authRoutes(api fiber.Router, db *gorm.DB) {
	adminRepository := repository.NewAdminRepository(db)
	authService := service.NewAuthService(adminRepository)
	authController := controller.NewAuthController(authService)

	api.Post("/admin/login", authController.Login)
}
