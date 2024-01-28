package dto

import (
	"go-absen/domain/user"
	"go-absen/entities"
	"strconv"
	"time"
)

type TGetUserResponse struct {
	ID       int       `json:"id"`
	Fullname string    `json:"fullname"`
	NIK      string    `json:"nik"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Gender   string    `json:"gender"`
	TimeNow  time.Time `json:"time_now"`
}

func GetUserResponse(user *entities.UserEntity) *TGetUserResponse {
	userFormatter := &TGetUserResponse{
		ID:       user.ID,
		Fullname: user.Fullname,
		NIK:      user.NIK,
		Phone:    user.Phone,
		Address:  user.Address,
		Gender:   user.Gender.Name,
		TimeNow:  time.Now(),
	}

	return userFormatter
}

type TUserResponse struct {
	ID        int    `json:"id"`
	Fullname  string `json:"fullname"`
	NIK       string `json:"nik"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Address   string `json:"address"`
	GenderID  int    `json:"gender_id"`
}

func GetUserLocationResponse(user *entities.UserEntity) *TUserResponse {
	return &TUserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		NIK:       user.NIK,
		Email:     user.Email,
		Phone:     user.Phone,
		Birthdate: user.Birthdate.Format("02-01-2006"), // Format sesuai dengan yang diminta
		Address:   user.Address,
		GenderID:  user.GenderID,
	}
}

type TRecordAttendanceRequest struct {
	Latitude  string `json:"latitude" validate:"required"`
	Longitude string `json:"longitude" validate:"required"`
}

// TAttendanceResponse adalah DTO untuk respons satu entitas absensi
type TAttendanceResponse struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	Latitude  float64        `json:"latitude"`
	Longitude float64        `json:"longitude"`
	CreatedAt time.Time      `json:"created_at"`
	User      *TUserResponse `json:"user"` // Change the type to *TGetUserResponse
}

func GetAttendanceResponse(attendance *entities.AttendanceEntity) *TAttendanceResponse {
	return &TAttendanceResponse{
		ID:        attendance.ID,
		UserID:    attendance.UserID,
		Latitude:  attendance.Latitude,
		Longitude: attendance.Longitude,
		CreatedAt: attendance.CreatedAt,
		User:      GetUserLocationResponse(&attendance.User),
	}
}

// TAttendanceHistoryResponse adalah DTO untuk respons riwayat absensi
type TAttendanceHistoryResponse []*TAttendanceResponse

// GetAttendanceHistoryResponse mengubah slice entitas absensi menjadi slice DTO respons
func GetAttendanceHistoryResponse(attendances []entities.AttendanceEntity) TAttendanceHistoryResponse {
	var responseList TAttendanceHistoryResponse
	for _, attendance := range attendances {
		responseList = append(responseList, GetAttendanceResponse(&attendance))
	}
	return responseList
}

type AttendanceHistoryResponse struct {
	Location  string `json:"location"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Status    string `json:"status"`
	Day       string `json:"day"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type AttendanceHistoryListResponse struct {
	Message string                      `json:"message"`
	Data    []AttendanceHistoryResponse `json:"data"`
}

func MapToAttendanceHistoryResponse(u user.UserServiceInterface, attendances []entities.AttendanceEntity) []AttendanceHistoryResponse {
	responseData := make([]AttendanceHistoryResponse, len(attendances))
	for i, a := range attendances {
		location, err := u.GetLocationName(a.Latitude, a.Longitude)
		if err != nil {
			// Handle error
			return nil
		}
		a.Location = location

		responseData[i] = AttendanceHistoryResponse{
			Location:  a.Location,
			Date:      a.CreatedAt.Format("02 January 2006"),
			Time:      a.CreatedAt.Format("15:04:05"),
			Status:    a.Status,
			Day:       a.CreatedAt.Format("Monday"),
			Latitude:  strconv.FormatFloat(a.Latitude, 'f', -1, 64),
			Longitude: strconv.FormatFloat(a.Longitude, 'f', -1, 64),
		}
	}
	return responseData
}