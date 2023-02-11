package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/TaichiOhmi/jwt-auth-with-goreact/app/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", middleware.Register)
	app.Post("/api/login", middleware.Login)
	app.Post("/api/logout", middleware.Logout)
	app.Get("/api/user", middleware.User)
}
