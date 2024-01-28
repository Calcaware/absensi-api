package dto

type TRegisterRequest struct {
	Fullname        string `json:"fullname" validate:"required"`
	NIK             string `json:"nik" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required"`
	Birthdate       string `json:"birthdate" validate:"required"`
	Address         string `json:"address" validate:"required"`
	GenderID        int    `json:"gender_id" validate:"required"`
	Password        string `json:"password" validate:"required,eqfield=PasswordConfirm"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

type TLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
