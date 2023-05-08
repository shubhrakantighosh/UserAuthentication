package requests

import "time"

type CreateUserServiceRequest struct {
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Mail         string    `json:"mail"`
	MobileNumber string    `json:"mobileNumber"`
	CreatedAt    time.Time `json:"createdAt"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    string    `json:"updatedBy"`
	DeletedAt    time.Time `json:"'deletedAt'"`
	DeletedBy    string    `json:"deletedBy"`
}

type UserLogInServiceRequest struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type ForgotUserRequest struct {
	Mail string `json:"mail"`
}

type HashPasswordRequest struct {
	Password string
}

type HashPasswordCheckRequest struct {
	Password string
}

type ChangeUserPasswordRequest struct {
	OldPassword string    `json:"old_password"`
	NewPassword string    `json:"new_password"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResetPasswordRequest struct {
	UserId      int       `json:"user_id"`
	UpdatedAt   time.Time `json:"updated_at"`
	NewPassword string    `json:"password"`
}

type ChangeUserPassword struct {
	Password  string
	UpdatedAt time.Time
}

type DeleteUserRequest struct {
	DeletedAt time.Time `json:"deleted_at"`
}
