package models

type LoginRequest struct {
	Username string `json:"username" goro:"unique" validate:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	AccessToken string `json:"access_token"`
}
