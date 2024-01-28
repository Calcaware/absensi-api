package auth

import (
	"github.com/gofiber/fiber/v2"
	"go-absen/domain/auth/dto"
	"go-absen/entities"
)

type AuthRepositoryInterface interface {
	InsertUser(newUser *entities.UserEntity) (*entities.UserEntity, error)
}

type AuthServiceInterface interface {
	Register(payload *dto.TRegisterRequest) (*entities.UserEntity, error)
	Login(payload *dto.TLoginRequest) (*entities.UserEntity, string, error)
}

type AuthHandlerInterface interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
