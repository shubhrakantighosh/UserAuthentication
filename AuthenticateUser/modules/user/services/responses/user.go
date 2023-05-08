package responses

type CreteUserResponse struct {
	Username string `json:"username"`
	Mail     string `json:"mail"`
}

type LoginUserResponse struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Mail     string   `json:"mail"`
	Meta     MetaData `json:"meta"`
}

type MetaData struct {
	AccessToken string `json:"access_token"`
}

type StatusMessage struct {
	Message string `json:"message"`
}

type ForgotPasswordServiceResponse struct {
	UserId int `json:"user_id"`
}

type HashPasswordResponse struct {
	Password string
}
