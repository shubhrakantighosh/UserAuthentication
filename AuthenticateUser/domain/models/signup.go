package models

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

type SignupResponse struct {
	Message string `json:"message"`
}
