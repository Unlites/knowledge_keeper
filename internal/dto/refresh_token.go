package dto

import validation "github.com/go-ozzo/ozzo-validation"

type RefreshTokenDTO struct {
	Token string `json:"refresh_token"`
}

func (r RefreshTokenDTO) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Token, validation.Required),
	)
}
