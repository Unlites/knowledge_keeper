package dto

import validation "github.com/go-ozzo/ozzo-validation"

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u UserDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required, validation.Length(2, 20)),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 16)),
	)
}

type TokensDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
