package router

import (
	"AuthenticateUser/controller/user"
	"AuthenticateUser/middleware"
	"github.com/gin-gonic/gin"
)

func Router(controller *user.Controller) *gin.Engine {
	router := gin.Default()
	users := router.Group("/users")
	{
		users.POST("/signup", controller.CreteUserController)
		users.POST("/login", controller.UserLogInController)
		users.POST("/logout", middleware.AuthenticateJWT, controller.LogoutUserController)
		users.POST("/change_password", middleware.AuthenticateJWT, controller.ChangeUserPassword)
		users.POST("/delete", middleware.AuthenticateJWT, controller.DeleteUserController)
		users.POST("/forgot_password", controller.ForgotUserController)
		users.POST("/rest_password", controller.ResetPasswordController)
		users.GET("/", controller.CountryCallingCodeByConditionController)
		users.GET("/search", controller.CountryCallingCodeByCharacterController)
	}
	return router
}
