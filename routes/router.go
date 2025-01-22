package internal

import (
	"fiber-poc-api/database/repository"
	"fiber-poc-api/handlers"
	"fiber-poc-api/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app fiber.Router, middleware fiber.Handler, db *gorm.DB) {

	// ==> Repositories
	userRepository := repository.NewUserRepository(db)
	loginHistoryRepository := repository.NewLoginHistoryRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	privilegeRepository := repository.NewPrivilegeRepository(db)
	rolePrivilegeRepository := repository.NewRolePrivilegeRepository(db)

	// ==> Services
	authService := services.NewAuthService(userRepository, loginHistoryRepository)
	roleService := services.NewRoleService(userRepository, roleRepository, privilegeRepository, rolePrivilegeRepository)

	// ==> Handlers
	authHandler := handlers.NewAuthHandler(authService)
	roleHandler := handlers.NewRoleHandler(roleService)

	app.Post("/api/v1/auth/login", authHandler.LoginHandler)
	app.Post("/api/v1/auth/register", authHandler.RegisterHandler)

	// middleware
	app.Get("/api/v1/user/get/all", middleware, authHandler.GetUserAllHandler)
	app.Get("/api/v1/role/create-role", middleware, roleHandler.CreateRoleHandler)
	app.Post("/api/v1/role/create-role-privilege", middleware, roleHandler.CreateRolePrivilegeHandler)

	// support
	app.Get("/api/v1/role/initial-permission", roleHandler.InitialHandler)
}
