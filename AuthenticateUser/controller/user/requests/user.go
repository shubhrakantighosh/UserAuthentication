package requests

type CreateUserRequest struct {
	Username     string `json:"username" validate:"required,alpha,min=5,max=20"`
	Password     string `json:"password" validate:"space,min=8,max=12,required,containsany=!@#$%^&*,uppercase,number"`
	Mail         string `json:"mail" validate:"required,email"`
	MobileNumber string `json:"mobile_number" validate:"required,numeric,len=10"`
}

type LogInUserRequest struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=8,max=12"`
}

type ForgotUserRequest struct {
	Mail string `json:"mail" validate:"required,email"`
}

type ChangeUserPasswordRequest struct {
	OldPassword        string `json:"old_password" validate:"required,min=8,max=12"`
	NewPassword        string `json:"new_password" validate:"required,min=8,max=12"`
	NewConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=12"`
}

type UpdateUserPasswordRequest struct {
	Password        string `json:"password" validate:"required,min=8,max=12"`
	ConfirmPassword string `json:"confirm_password" validate:"required,required,min=8,max=12"`
}

type ResetPasswordRequest struct {
	UserId          int    `json:"user_id" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
