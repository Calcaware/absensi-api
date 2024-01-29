package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-absen/domain/user"
	"go-absen/domain/user/dto"
	"go-absen/entities"
	"go-absen/helper/cloudinary"
	"go-absen/helper/response"
	"mime/multipart"
	"strconv"
)

type userHandler struct {
	userService user.UserServiceInterface
}

func NewUserHandler(userService user.UserServiceInterface) user.UserHandlerInterface {
	return &userHandler{userService: userService}
}

func (u userHandler) GetMe(c *fiber.Ctx) error {
	user, ok := c.Locals("CurrentUser").(*entities.UserEntity)
	if !ok || user == nil {
		return response.SendStatusUnauthorized(c, "user not found")
	}

	return response.GetCurrentUser(c, dto.GetUserResponse(user))
}

func (u userHandler) RecordAttendance(c *fiber.Ctx) error {
	userEntity, ok := c.Locals("CurrentUser").(*entities.UserEntity)
	if !ok || userEntity == nil {
		return response.SendStatusUnauthorized(c, "user not found")
	}
	payload := new(dto.TRecordAttendanceRequest)
	if err := c.BodyParser(payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid request payload")
	}
	latitude, err := strconv.ParseFloat(payload.Latitude, 64)
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid latitude")
	}
	longitude, err := strconv.ParseFloat(payload.Longitude, 64)
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid longitude")
	}

	_, err = u.userService.RecordAttendance(userEntity.ID, latitude, longitude)
	if err != nil {
		return response.SendStatusInternalServerError(c, "failed to record attendance"+err.Error())
	}
	return response.SendStatusOkResponse(c, "success record attendance")
}

func (u userHandler) GetAttendanceHistory(c *fiber.Ctx) error {
	userEntity, ok := c.Locals("CurrentUser").(*entities.UserEntity)
	if !ok || userEntity == nil {
		return response.SendStatusUnauthorized(c, "user not found")
	}

	attendances, err := u.userService.GetAttendanceHistory(userEntity.ID)
	if err != nil {
		return response.SendStatusInternalServerError(c, "failed to get attendance history")
	}

	responseData := dto.MapToAttendanceHistoryResponse(u.userService, attendances)

	responseStruct := dto.AttendanceHistoryListResponse{
		Message: "Attendance History",
		Data:    responseData,
	}

	return response.SendStatusOkWithDataResponse(c, responseStruct.Message, responseStruct.Data)
}

func (u userHandler) UpdateAvatar(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("CurrentUser").(*entities.UserEntity)
	if !ok || currentUser == nil {
		return response.SendStatusUnauthorized(c, "User not found")
	}

	file, err := c.FormFile("avatar")
	var uploadedURL string
	if err == nil {
		fileToUpload, err := file.Open()
		if err != nil {
			return response.SendStatusInternalServerError(c, "Failed open to open file: "+err.Error())
		}
		defer func(fileToUpload multipart.File) {
			_ = fileToUpload.Close()
		}(fileToUpload)
		uploadedURL, err = cloudinary.Uploader(fileToUpload)
		if err != nil {
			return response.SendStatusInternalServerError(c, "Failed to upload image: "+err.Error())
		}
	}

	image, err := u.userService.UpdateAvatar(currentUser.ID, uploadedURL)
	if err != nil {
		return response.SendStatusBadRequest(c, "Error upload image: "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Success updating avatar", dto.UpdateAvatarResponse(image))
}
