package dto

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokensDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
