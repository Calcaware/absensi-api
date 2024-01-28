package service

import (
	"errors"
	"go-absen/domain/auth"
	"go-absen/domain/auth/dto"
	"go-absen/domain/user"
	"go-absen/entities"
	"go-absen/helper/hashing"
	"go-absen/helper/jwt"
	"time"
)

type AuthService struct {
	repo        auth.AuthRepositoryInterface
	userService user.UserServiceInterface
	hashing     hashing.HashInterface
	jwt         jwt.IJwt
}

func NewAuthService(repo auth.AuthRepositoryInterface, userService user.UserServiceInterface, hashing hashing.HashInterface, jwt jwt.IJwt) auth.AuthServiceInterface {
	return &AuthService{repo: repo, userService: userService, hashing: hashing, jwt: jwt}
}

func (s *AuthService) Register(payload *dto.TRegisterRequest) (*entities.UserEntity, error) {
	isExistEmail, _ := s.userService.GetEmail(payload.Email)
	if isExistEmail != nil {
		return nil, errors.New("your email has been already")
	}

	isExistNik, _ := s.userService.GetNik(payload.NIK)
	if isExistNik != nil {
		return nil, errors.New("your nik has been already")
	}

	if payload.Password != payload.PasswordConfirm {
		return nil, errors.New("password does not match")
	}

	hashPassword, err := s.hashing.GenerateHash(payload.Password)
	if err != nil {
		return nil, err
	}

	birthdate, err := time.Parse("02-01-2006", payload.Birthdate)
	if err != nil {
		return nil, errors.New("error parsing birthdate")
	}

	createdAtUnix := time.Now().Unix()
	newUser := &entities.UserEntity{
		Fullname:  payload.Fullname,
		NIK:       payload.NIK,
		Email:     payload.Email,
		Phone:     payload.Phone,
		Birthdate: birthdate,
		Address:   payload.Address,
		GenderID:  payload.GenderID,
		Password:  hashPassword,
		CreatedAt: createdAtUnix,
	}

	user, err := s.repo.InsertUser(newUser)
	if err != nil {
		return nil, errors.New("error inserting user")
	}

	return user, nil
}

func (s *AuthService) Login(payload *dto.TLoginRequest) (*entities.UserEntity, string, error) {
	user, err := s.userService.GetEmail(payload.Email)
	if err != nil {
		return nil, "", errors.New("email not found")
	}

	isValidPassword, err := s.hashing.ComparePassword(user.Password, payload.Password)
	if err != nil || !isValidPassword {
		return nil, "", errors.New("incorrect password")
	}

	accessSecret, err := s.jwt.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, "", err
	}

	return user, accessSecret, nil
}
