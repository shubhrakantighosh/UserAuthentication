package services

import (
	"AuthenticateUser/domain/interfaces"
	"AuthenticateUser/domain/models"
	"AuthenticateUser/errors"
	_ "AuthenticateUser/initializers"
	"AuthenticateUser/modules/user/services/adapters"
	"AuthenticateUser/modules/user/services/requests"
	"AuthenticateUser/modules/user/services/responses"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserService struct {
	repository interfaces.UserRepository
}

func NewService(userRepository interfaces.UserRepository) *UserService {
	return &UserService{
		repository: userRepository,
	}
}

func (svc *UserService) CreteUserService(request *requests.CreateUserServiceRequest) (*responses.CreteUserResponse, errors.CustomError) {
	result := adapters.SvcToRepoCreateUser(request)
	res, cusErr := hashPassword(requests.HashPasswordRequest{Password: request.Password})
	if cusErr.Exist() {
		return nil, cusErr
	}
	result.Password = res.Password

	response, cusErr := svc.repository.CreateUserRepository(result)
	if cusErr.ErrorExist {
		return nil, cusErr
	}
	return response, errors.CustomError{}
}

func (svc *UserService) UserLogInService(request *requests.UserLogInServiceRequest) (*responses.LoginUserResponse, errors.CustomError) {
	result := adapters.SvcToRepoLoginUser(request)

	username := map[string]interface{}{
		"username":   result.Username,
		"deleted_at": nil,
	}
	user, cusErr := svc.repository.GetUserByConditions(username)
	if cusErr.Exist() {
		return nil, cusErr
	}

	cusErr = hashPasswordCheck(user.Password, result.Password)
	if cusErr.Exist() {
		return nil, cusErr
	}

	userSessionByCondition := map[string]interface{}{
		"user_id":  user.Id,
		"end_time": nil,
	}

	_, cusErr = svc.repository.GetUserSessionByConditions(userSessionByCondition)
	if !cusErr.Exist() {
		return nil, errors.CustomError{Time: time.Now(), UserMessage: "please contact admin", ErrorExist: true}
	}

	createUserSession := adapters.SvcToRepoCreateUserSession(result.CreatedAt, user.Id)
	cusErr = svc.repository.CreateUserSession(createUserSession)
	if cusErr.Exist() {
		return nil, cusErr
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	secretKey := os.Getenv("SecretKey")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, errors.CustomError{Time: time.Now(), UserMessage: err.Error(), ErrorExist: true}
	}

	resp := adapters.SvcToCtrlLoginUserResp(user, tokenString)
	return resp, errors.CustomError{}
}

func (svc *UserService) LogoutUserService(userId int) (*responses.StatusMessage, errors.CustomError) {
	userEndTime := map[string]interface{}{
		"user_id":  userId,
		"end_time": nil,
	}

	userSession, cusErr := svc.repository.GetUserSessionByConditions(userEndTime)
	if cusErr.Exist() {
		return nil, cusErr
	}

	userSessionId := map[string]interface{}{
		"id": userSession.Id,
	}

	userID := map[string]interface{}{
		"id":         userId,
		"deleted_at": nil,
	}
	if _, cusErr := svc.repository.GetUserByConditions(userID); cusErr.Exist() {
		return nil, cusErr
	}

	cusErr = svc.repository.EndUserSession(userSessionId)
	if cusErr.Exist() {
		return nil, cusErr
	}

	return &responses.StatusMessage{Message: "Logout successfully."}, errors.CustomError{}
}

func (svc *UserService) ForgotUserService(request *requests.ForgotUserRequest) (*responses.ForgotPasswordServiceResponse, errors.CustomError) {
	userMail := map[string]interface{}{
		"mail": request.Mail,
	}
	user, cusErr := svc.repository.GetUserByConditions(userMail)
	if cusErr.Exist() {
		return nil, cusErr
	}

	result := adapters.SvcToCtrlForgotUserResp(user.Id)
	return result, errors.CustomError{}
}

func (svc *UserService) UserPasswordChangeService(request *requests.ChangeUserPasswordRequest, userId int) (*responses.StatusMessage, errors.CustomError) {

	userEndTime := map[string]interface{}{
		"user_id":  userId,
		"end_time": nil,
	}

	_, cusErr := svc.repository.GetUserSessionByConditions(userEndTime)
	if cusErr.Exist() {
		return nil, cusErr
	}

	userID := map[string]interface{}{
		"id":         userId,
		"deleted_at": nil,
	}
	user, cusErr := svc.repository.GetUserByConditions(userID)
	if cusErr.Exist() {
		return nil, cusErr
	}

	cusErr = hashPasswordCheck(user.Password, request.OldPassword)
	if cusErr.Exist() {
		return nil, cusErr
	}

	newHashPassword, cusErr := hashPassword(requests.HashPasswordRequest{Password: request.NewPassword})
	if cusErr.Exist() {
		return nil, cusErr
	}

	result := adapters.SvcToRepUserPasswordChangeRequest(newHashPassword.Password, request.UpdatedAt)
	cusErr = svc.repository.UserPasswordChangeRepository(userID, result)
	if cusErr.Exist() {
		return nil, cusErr
	}
	return &responses.StatusMessage{Message: "password updated successfully."}, errors.CustomError{}
}

func (svc *UserService) DeleteUserService(userId int, request *requests.DeleteUserRequest) (*responses.StatusMessage, errors.CustomError) {
	userEndTime := map[string]interface{}{
		"user_id":  userId,
		"end_time": nil,
	}

	userSession, cusErr := svc.repository.GetUserSessionByConditions(userEndTime)
	if cusErr.Exist() {
		return nil, cusErr
	}

	userSessionId := map[string]interface{}{
		"id": userSession.Id,
	}

	userID := map[string]interface{}{
		"id":         userId,
		"deleted_at": nil,
	}
	user, cusErr := svc.repository.GetUserByConditions(userID)
	if cusErr.Exist() {
		return nil, cusErr
	}

	cusErr = svc.repository.EndUserSession(userSessionId)
	if cusErr.Exist() {
		return nil, cusErr
	}

	result := adapters.SvcToRepoDeleteUserRequest(user.Username, request)
	cusErr = svc.repository.DeleteUserRepository(userID, result)
	if cusErr.Exist() {
		return nil, cusErr
	}
	return &responses.StatusMessage{Message: "Deleted successfully."}, errors.CustomError{}
}

func (svc *UserService) ResetPasswordController(request *requests.ResetPasswordRequest) (*responses.StatusMessage, errors.CustomError) {
	userId := map[string]interface{}{
		"id":         request.UserId,
		"deleted_at": nil,
	}

	result := adapters.SvcToRepUserPasswordChangeRequest(request.NewPassword, request.UpdatedAt)
	cusErr := svc.repository.UserPasswordChangeRepository(userId, result)
	if cusErr.Exist() {
		return nil, cusErr
	}

	return &responses.StatusMessage{Message: "Password Updated successfully."}, errors.CustomError{}
}

func hashPassword(request requests.HashPasswordRequest) (*responses.HashPasswordResponse, errors.CustomError) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.CustomError{Time: time.Now(), UserMessage: "something went wrong", Error: err, ErrorExist: true}
	}
	result := string(hashedPassword)
	return &responses.HashPasswordResponse{Password: result}, errors.CustomError{}
}

func hashPasswordCheck(user string, result string) errors.CustomError {
	err := bcrypt.CompareHashAndPassword([]byte(user), []byte(result))
	if err != nil {
		return errors.CustomError{Time: time.Now(), Error: err, UserMessage: "Wrong Password", ErrorExist: true}
	}
	return errors.CustomError{}
}

func (svc *UserService) CountryCallingCodeByConditionService(filter map[string]interface{}) ([]*models.CountriesCallingCode, errors.CustomError) {
	response, cusErr := svc.repository.CountryCallingCodeByConditionRepository(filter)
	if cusErr.Exist() {
		return nil, cusErr
	}
	return response, errors.CustomError{}
}

func (svc *UserService) CountryCallingCodeByCharacterService(filter map[string]string) ([]*models.CountriesCallingCode, errors.CustomError) {
	value := ""
	for k, v := range filter {
		if len(k) == 0 {
			response, cusErr := svc.repository.CountryCallingCodeByConditionRepository(map[string]interface{}{})
			if cusErr.Exist() {
				return nil, cusErr
			}
			return response, errors.CustomError{}
		}
		value = v
	}
	response, cusErr := svc.repository.CountryCallingCodeByCharacterRepository(value)
	if cusErr.Exist() {
		return nil, cusErr
	}
	return response, errors.CustomError{}
}
