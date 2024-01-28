package user

import (
	"github.com/gofiber/fiber/v2"
	"go-absen/entities"
	"net/http"
)

type UserRepositoryInterface interface {
	FindId(id int) (*entities.UserEntity, error)
	FindEmail(email string) (*entities.UserEntity, error)
	FindNik(nik string) (*entities.UserEntity, error)
	InsertAttendance(attendance *entities.AttendanceEntity) (*entities.AttendanceEntity, error)
	GetAttendanceHistory(userID int) ([]entities.AttendanceEntity, error)
	GetAttendanceByDate(userID int, date string) (*entities.AttendanceEntity, error)
}

type UserServiceInterface interface {
	GetId(id int) (*entities.UserEntity, error)
	GetEmail(email string) (*entities.UserEntity, error)
	GetNik(nik string) (*entities.UserEntity, error)
	RecordAttendance(userID int, latitude, longitude float64) (*entities.AttendanceEntity, error)
	GetAttendanceHistory(userID int) ([]entities.AttendanceEntity, error)
	GetLocationName(latitude, longitude float64) (string, error)
	UnmarshalResponseBody(response *http.Response, v interface{}) error
}

type UserHandlerInterface interface {
	GetMe(c *fiber.Ctx) error
	RecordAttendance(c *fiber.Ctx) error
	GetAttendanceHistory(c *fiber.Ctx) error
}
