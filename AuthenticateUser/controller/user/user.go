package user

import (
	"AuthenticateUser/controller/constants"
	"AuthenticateUser/controller/user/adapters"
	"AuthenticateUser/controller/user/requests"
	"AuthenticateUser/controller/utils"
	"AuthenticateUser/controller/validate"
	"AuthenticateUser/domain/interfaces"
	"AuthenticateUser/modules/user/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Controller struct {
	service interfaces.UserService
}

func NewController(service interfaces.UserService) *Controller {
	return &Controller{
		service: service,
	}
}

func (wc *Controller) CreteUserController(ctx *gin.Context) {
	var request *requests.CreateUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	createUserRequest := trimCreteUserController(request)
	cusErr := validate.Get(createUserRequest)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	result := adapters.CtrlToSvcCreateUserRequest(createUserRequest)
	response, cusErr := wc.service.CreteUserService(result)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}

func trimCreteUserController(request *requests.CreateUserRequest) *requests.CreateUserRequest {
	return &requests.CreateUserRequest{
		Username:     strings.TrimSpace(request.Username),
		Password:     strings.TrimSpace(request.Password),
		Mail:         strings.TrimSpace(request.Mail),
		MobileNumber: strings.TrimSpace(request.MobileNumber),
	}
}

func (wc *Controller) UserLogInController(ctx *gin.Context) {
	var request *requests.LogInUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	logInUserRequest := trimUserLogInController(request)

	cusErr := validate.Get(logInUserRequest)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	result := adapters.CtrlToSvcUserLogInRequest(logInUserRequest)
	response, cusErr := wc.service.UserLogInService(result)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", response.Meta.AccessToken, 3600, "", "", false, true)

	ctx.JSON(http.StatusOK, response)
	return
}

func trimUserLogInController(request *requests.LogInUserRequest) *requests.LogInUserRequest {
	return &requests.LogInUserRequest{
		Username: strings.TrimSpace(request.Username),
		Password: strings.TrimSpace(request.Password),
	}
}

func (wc *Controller) LogoutUserController(ctx *gin.Context) {
	userId, cusErr := utils.GetUserId(ctx)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	message, cusErr := wc.service.LogoutUserService(userId)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}
	ctx.JSON(http.StatusOK, message)
	return
}

func (wc *Controller) ForgotUserController(ctx *gin.Context) {
	var forgotUserRequest *requests.ForgotUserRequest
	err := ctx.ShouldBindJSON(&forgotUserRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	cusErr := validate.Get(forgotUserRequest)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	result := adapters.CtrlToSvcForgotUserRequest(forgotUserRequest)
	response, cusErr := wc.service.ForgotUserService(result)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}

func (wc *Controller) ChangeUserPassword(ctx *gin.Context) {
	userId, cusErr := utils.GetUserId(ctx)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	var request *requests.ChangeUserPasswordRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	changeUserPasswordRequest := trimChangeUserPassword(request)

	cusErr = validate.Get(changeUserPasswordRequest)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	if changeUserPasswordRequest.NewPassword != changeUserPasswordRequest.NewConfirmPassword {
		ctx.JSON(http.StatusBadRequest, "Error : password mismatch")
		return
	}

	result := adapters.CtrlToSvcUserPasswordChangeRequest(changeUserPasswordRequest)
	response, cusErr := wc.service.UserPasswordChangeService(result, userId)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func trimChangeUserPassword(request *requests.ChangeUserPasswordRequest) *requests.ChangeUserPasswordRequest {
	return &requests.ChangeUserPasswordRequest{
		OldPassword:        strings.TrimSpace(request.OldPassword),
		NewPassword:        strings.TrimSpace(request.NewPassword),
		NewConfirmPassword: strings.TrimSpace(request.NewConfirmPassword),
	}
}

func (wc *Controller) DeleteUserController(ctx *gin.Context) {
	userId, cusErr := utils.GetUserId(ctx)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	result := adapters.CtrlToSvcDeleteUserRequest()
	response, cusErr := wc.service.DeleteUserService(userId, result)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (wc *Controller) ResetPasswordController(ctx *gin.Context) {
	var resetPasswordRequest *requests.ResetPasswordRequest
	err := ctx.ShouldBindJSON(&resetPasswordRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	cusErr := validate.Get(resetPasswordRequest)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}

	result := adapters.CtrlToSvcUserRestPasswordRequest(resetPasswordRequest)
	response, cusErr := wc.service.ResetPasswordController(result)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
	return
}

func (wc *Controller) SearchProducts(ctx *gin.Context) {
	productIds := ctx.Query(constants.SearchByProductId)
	filter := map[string][]string{}
	if len(productIds) > 0 {
		filter[constants.SearchByProductId] = []string{productIds}
	}

	res := repository.SearchProducts(filter)
	ctx.JSON(http.StatusOK, res)
	return
}

func (wc *Controller) CountryCallingCodeByConditionController(ctx *gin.Context) {
	filter := map[string]interface{}{}
	countryName := ctx.Query(constants.CountryName)

	if len(countryName) > 0 {
		filter[constants.CountryName] = countryName
	}

	res, cusErr := wc.service.CountryCallingCodeByConditionService(filter)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (wc *Controller) CountryCallingCodeByCharacterController(ctx *gin.Context) {
	filter := map[string]string{}
	countryNameLike := ctx.Query(constants.CountryNameSearch)
	if len(countryNameLike) > 0 {
		filter[constants.CountryName] = countryNameLike
	}

	response, cusErr := wc.service.CountryCallingCodeByCharacterService(filter)
	if cusErr.Exist() {
		ctx.JSON(http.StatusBadRequest, cusErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
