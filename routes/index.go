package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-absen/domain/auth"
	"go-absen/domain/user"
	"go-absen/helper/jwt"
	"go-absen/middleware/authentication"
)

func BootAuthRoute(app *fiber.App, handler auth.AuthHandlerInterface) {
	authGroup := app.Group("api/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
}

func BootUserRoute(app *fiber.App, handler user.UserHandlerInterface, jwtService jwt.IJwt, userService user.UserServiceInterface) {
	userGroup := app.Group("api/user")
	userGroup.Get("/me", authentication.Protected(jwtService, userService), handler.GetMe)
	userGroup.Post("/attendance", authentication.Protected(jwtService, userService), handler.RecordAttendance)
	userGroup.Get("/attendance", authentication.Protected(jwtService, userService), handler.GetAttendanceHistory)
}
