package adapters

import (
	"AuthenticateUser/controller/user/requests"
	svc "AuthenticateUser/modules/user/services/requests"
	"time"
)

func CtrlToSvcCreateUserRequest(request *requests.CreateUserRequest) *svc.CreateUserServiceRequest {
	return &svc.CreateUserServiceRequest{
		Username:     request.Username,
		Password:     request.Password,
		Mail:         request.Mail,
		MobileNumber: request.MobileNumber,
		CreatedAt:    time.Now(),
		CreatedBy:    request.Username,
		UpdatedAt:    time.Now(),
		UpdatedBy:    request.Username,
	}
}

func CtrlToSvcUserLogInRequest(request *requests.LogInUserRequest) *svc.UserLogInServiceRequest {
	return &svc.UserLogInServiceRequest{
		Username:  request.Username,
		Password:  request.Password,
		CreatedAt: time.Now(),
	}
}

func CtrlToSvcForgotUserRequest(request *requests.ForgotUserRequest) *svc.ForgotUserRequest {
	return &svc.ForgotUserRequest{
		Mail: request.Mail,
	}
}

func CtrlToSvcUserPasswordChangeRequest(request *requests.ChangeUserPasswordRequest) *svc.ChangeUserPasswordRequest {
	return &svc.ChangeUserPasswordRequest{
		OldPassword: request.OldPassword,
		NewPassword: request.NewPassword,
		UpdatedAt:   time.Now(),
	}
}

func CtrlToSvcUserRestPasswordRequest(request *requests.ResetPasswordRequest) *svc.ResetPasswordRequest {
	return &svc.ResetPasswordRequest{
		UserId:      request.UserId,
		NewPassword: request.Password,
		UpdatedAt:   time.Now(),
	}
}

func CtrlToSvcDeleteUserRequest() *svc.DeleteUserRequest {
	return &svc.DeleteUserRequest{
		DeletedAt: time.Now(),
	}
}
