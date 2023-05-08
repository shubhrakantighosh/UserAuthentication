package interfaces

import (
	"AuthenticateUser/domain/models"
	"AuthenticateUser/errors"
	"AuthenticateUser/modules/user/services/requests"
	"AuthenticateUser/modules/user/services/responses"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreteUserController(*gin.Context)
	UserLogInController(*gin.Context)
	ForgotUserController(*gin.Context)
	ChangeUserPassword(*gin.Context)
	LogoutUserController(*gin.Context)
	DeleteUserController(*gin.Context)
	ResetPasswordController(*gin.Context)
	CountryCallingCodeByConditionController(*gin.Context)
	CountryCallingCodeByCharacterController(*gin.Context)
}

type UserService interface {
	CreteUserService(*requests.CreateUserServiceRequest) (*responses.CreteUserResponse, errors.CustomError)
	UserLogInService(*requests.UserLogInServiceRequest) (*responses.LoginUserResponse, errors.CustomError)
	LogoutUserService(int) (*responses.StatusMessage, errors.CustomError)
	ForgotUserService(*requests.ForgotUserRequest) (*responses.ForgotPasswordServiceResponse, errors.CustomError)
	UserPasswordChangeService(*requests.ChangeUserPasswordRequest, int) (*responses.StatusMessage, errors.CustomError)
	DeleteUserService(int, *requests.DeleteUserRequest) (*responses.StatusMessage, errors.CustomError)
	ResetPasswordController(*requests.ResetPasswordRequest) (*responses.StatusMessage, errors.CustomError)
	CountryCallingCodeByConditionService(map[string]interface{}) ([]*models.CountriesCallingCode, errors.CustomError)
	CountryCallingCodeByCharacterService(map[string]string) ([]*models.CountriesCallingCode, errors.CustomError)
}

type UserRepository interface {
	CreateUserRepository(*models.Users) (*responses.CreteUserResponse, errors.CustomError)
	GetUserByConditions(map[string]interface{}) (*models.Users, errors.CustomError)
	CreateUserSession(*models.UserSessions) errors.CustomError
	GetUserSessionByConditions(map[string]interface{}) (*models.UserSessions, errors.CustomError)
	EndUserSession(map[string]interface{}) errors.CustomError
	UserPasswordChangeRepository(map[string]interface{}, *models.Users) errors.CustomError
	DeleteUserRepository(map[string]interface{}, *models.Users) errors.CustomError
	CountryCallingCodeByConditionRepository(map[string]interface{}) ([]*models.CountriesCallingCode, errors.CustomError)
	CountryCallingCodeByCharacterRepository(string) ([]*models.CountriesCallingCode, errors.CustomError)
}
