package user

import (
	"AuthenticateUser/modules/user/repository"
	"AuthenticateUser/modules/user/services"
	"gorm.io/gorm"
)

func Wire(db *gorm.DB) *Controller {
	repo := repository.NewRepository(db)
	svc := services.NewService(repo)
	ctrl := NewController(svc)
	return ctrl
}
