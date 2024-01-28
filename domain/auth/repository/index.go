package repository

import (
	"go-absen/domain/auth"
	"go-absen/entities"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.AuthRepositoryInterface {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) InsertUser(newUser *entities.UserEntity) (*entities.UserEntity, error) {
	if err := r.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}
