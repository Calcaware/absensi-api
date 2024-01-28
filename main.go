package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-absen/config"
	rUser "go-absen/domain/user/repository"
	sUser "go-absen/domain/user/service"
	"go-absen/helper/database"
	"go-absen/helper/hashing"
	jwt2 "go-absen/helper/jwt"
	"go-absen/middleware/logging"
	"go-absen/routes"

	hUser "go-absen/domain/user/handler"

	hAuth "go-absen/domain/auth/handler"
	rAuth "go-absen/domain/auth/repository"
	sAuth "go-absen/domain/auth/service"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Welcome to API Absen",
		CaseSensitive: false,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Authorization",
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))

	var bootConfig = config.BootConfig()
	db := database.BootDatabase(*bootConfig)
	database.MigrateTable(db)
	hash := hashing.NewHash()
	jwt := jwt2.NewJWT(bootConfig.SecretKey)

	userRepo := rUser.NewUserRepository(db)
	userService := sUser.NewUserService(userRepo, hash)
	userHandler := hUser.NewUserHandler(userService)

	authRepo := rAuth.NewAuthRepository(db)
	authService := sAuth.NewAuthService(authRepo, userService, hash, jwt)
	authHandler := hAuth.NewAuthHandler(authService)

	app.Use(logging.Logging())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Ping Ping",
		})
	})

	routes.BootAuthRoute(app, authHandler)
	routes.BootUserRoute(app, userHandler, jwt, userService)

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)
	if err := app.Listen(addr).Error(); err != addr {
		panic("Appilaction failed to start")
	}
}
