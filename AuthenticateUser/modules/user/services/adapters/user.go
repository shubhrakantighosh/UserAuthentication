package adapters

import (
	"AuthenticateUser/domain/models"
	svc "AuthenticateUser/modules/user/services/requests"
	"AuthenticateUser/modules/user/services/responses"
	"gopkg.in/guregu/null.v4"
	"time"
)

func SvcToRepoCreateUser(request *svc.CreateUserServiceRequest) *models.Users {
	return &models.Users{
		Username:     request.Username,
		Password:     request.Password,
		Mail:         request.Mail,
		MobileNumber: request.MobileNumber,
		CreatedAt:    null.TimeFrom(request.CreatedAt),
		CreatedBy:    null.StringFrom(request.Username),
		UpdatedAt:    null.TimeFrom(request.UpdatedAt),
		UpdatedBy:    null.StringFrom(request.UpdatedBy),
	}
}

func SvcToRepoLoginUser(request *svc.UserLogInServiceRequest) *models.Users {
	return &models.Users{
		Username:  request.Username,
		Password:  request.Password,
		CreatedAt: null.TimeFrom(request.CreatedAt),
	}
}

func SvcToRepoCreateUserSession(startTime null.Time, userId int) *models.UserSessions {
	return &models.UserSessions{
		StartTime: startTime,
		UserId:    userId,
	}
}

func SvcToCtrlLoginUserResp(request *models.Users, token string) *responses.LoginUserResponse {
	return &responses.LoginUserResponse{
		Id:       request.Id,
		Username: request.Username,
		Mail:     request.Mail,
		Meta: responses.MetaData{
			AccessToken: token,
		},
	}
}

func SvcToCtrlForgotUserResp(userId int) *responses.ForgotPasswordServiceResponse {
	return &responses.ForgotPasswordServiceResponse{
		UserId: userId,
	}
}

func SvcToRepUserPasswordChangeRequest(password string, updatedAt time.Time) *models.Users {
	return &models.Users{
		Password:  password,
		UpdatedAt: null.TimeFrom(updatedAt),
	}
}

func SvcToRepoDeleteUserRequest(username string, request *svc.DeleteUserRequest) *models.Users {
	return &models.Users{
		DeletedAt: null.TimeFrom(request.DeletedAt),
		DeletedBy: null.StringFrom(username),
	}
}
