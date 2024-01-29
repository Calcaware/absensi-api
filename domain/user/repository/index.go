package repository

import (
	"fmt"
	"go-absen/domain/user"
	"go-absen/entities"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindId(id int) (*entities.UserEntity, error) {
	var user *entities.UserEntity
	if err := r.db.Preload("Gender").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindEmail(email string) (*entities.UserEntity, error) {
	var user *entities.UserEntity
	if err := r.db.Table("users").
		Where("email = ? AND deleted_at IS NULL", email).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindNik(nik string) (*entities.UserEntity, error) {
	var user *entities.UserEntity
	if err := r.db.Table("users").
		Where("nik = ? AND deleted_at IS NULL", nik).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) InsertAttendance(attendance *entities.AttendanceEntity) (*entities.AttendanceEntity, error) {
	result := r.db.Create(attendance)
	if result.Error != nil {
		return nil, result.Error
	}

	return attendance, nil
}

func (r *UserRepository) GetAttendanceHistory(userID int) ([]entities.AttendanceEntity, error) {
	var attendances []entities.AttendanceEntity
	result := r.db.Where("user_id = ?", userID).Order("created_at desc").Limit(7).Find(&attendances)
	if result.Error != nil {
		return nil, result.Error
	}

	return attendances, nil
}

func (r *UserRepository) GetAttendanceByDate(userID int, date string) (*entities.AttendanceEntity, error) {
	var attendance entities.AttendanceEntity
	startOfDay, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	endOfDay := startOfDay.Add(24 * time.Hour)

	// Gunakan zona waktu "Asia/Jakarta" pada saat pencarian
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	fmt.Printf("Searching for attendance between %v and %v\n", startOfDay.In(location), endOfDay.In(location))

	if err := r.db.
		Where("user_id = ? AND created_at >= ? AND created_at < ?", userID, startOfDay.In(location).Unix(), endOfDay.In(location).Unix()).
		First(&attendance).
		Error; err != nil {
		return nil, err
	}

	fmt.Printf("Found attendance at %v\n", time.Unix(attendance.CreatedAt, 0).In(location))

	return &attendance, nil
}

func (r *UserRepository) UpdateUserAvatar(userID int, avatarPath string) error {
	user := &entities.UserEntity{ID: userID}

	result := r.db.Model(user).Updates(map[string]interface{}{"avatar": avatarPath})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
