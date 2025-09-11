package router

import (
	"app/src/controller"
	m "app/src/middleware"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(v1 fiber.Router, u service.UserService, t service.TokenService) {
	userController := controller.NewUserController(u, t)

	user := v1.Group("/users")

	user.Get("/me", m.Auth(u), userController.GetMe)
	user.Get("/", m.Auth(u), userController.GetUsers)
	user.Post("/", m.Auth(u), userController.CreateUser)
	user.Get("/:userId", m.Auth(u), userController.GetUserByID)
	user.Patch("/:userId", m.Auth(u), userController.UpdateUser)
	user.Delete("/:userId", m.Auth(u), userController.DeleteUser)

}
