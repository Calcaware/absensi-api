package dto

import "go-absen/entities"

type TLoginResponse struct {
	Avatar   string `json:"avatar"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"access_token"`
}

func LoginResponse(user *entities.UserEntity, token string) *TLoginResponse {
	userFormatter := &TLoginResponse{}
	userFormatter.Avatar = user.Avatar
	userFormatter.Fullname = user.Fullname
	userFormatter.Email = user.Email
	userFormatter.Token = token

	return userFormatter
}
