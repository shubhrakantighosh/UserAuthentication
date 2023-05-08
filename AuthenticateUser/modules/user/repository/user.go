package repository

import (
	"AuthenticateUser/domain/models"
	"AuthenticateUser/errors"
	"AuthenticateUser/modules/user/services/responses"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	repository *gorm.DB
}

func NewRepository(repository *gorm.DB) *Repository {
	return &Repository{
		repository: repository,
	}
}

func (repo *Repository) CreateUserRepository(request *models.Users) (*responses.CreteUserResponse, errors.CustomError) {
	err := repo.repository.Model(models.Users{}).Create(&request).Error
	if err != nil {
		cusErr := errors.DbError(err)
		return nil, cusErr
	}
	return &responses.CreteUserResponse{Username: request.Username, Mail: request.Mail}, errors.CustomError{}
}

func (repo *Repository) GetUserSessionByConditions(filter map[string]interface{}) (res *models.UserSessions, cusErr errors.CustomError) {
	err := repo.repository.Model(&models.UserSessions{}).Where(filter).First(&res).Error
	if err != nil {
		message := map[string]string{}
		message["Error"] = "Please login first."
		return nil, errors.CustomError{Time: time.Now(), StructErrors: message, ErrorExist: true}
	}
	return res, errors.CustomError{}
}

func (repo *Repository) EndUserSession(filter map[string]interface{}) errors.CustomError {
	err := repo.repository.Model(&models.UserSessions{}).Where(filter).Update("end_time", time.Now()).Error
	if err != nil {
		cusErr := errors.DbError(err)
		return cusErr
	}
	return errors.CustomError{}
}

func (repo *Repository) UserPasswordChangeRepository(filter map[string]interface{}, request *models.Users) errors.CustomError {
	err := repo.repository.Model(&models.Users{}).Where(filter).Updates(request).Error
	if err != nil {
		cusErr := errors.DbError(err)
		return cusErr
	}
	return errors.CustomError{}
}

func (repo *Repository) GetUserByConditions(filter map[string]interface{}) (res *models.Users, cusErr errors.CustomError) {
	err := repo.repository.Model(&models.Users{}).Where(filter).First(&res).Error
	if err != nil {
		cusErr = errors.DbError(err)
		return nil, cusErr
	}
	return res, errors.CustomError{}
}

func (repo *Repository) CreateUserSession(request *models.UserSessions) errors.CustomError {
	err := repo.repository.Model(&models.UserSessions{}).Create(&request).Error
	if err != nil {
		cusErr := errors.DbError(err)
		return cusErr
	}
	return errors.CustomError{}
}

func (repo *Repository) DeleteUserRepository(filter map[string]interface{}, request *models.Users) errors.CustomError {
	err := repo.repository.Model(&models.Users{}).Where(filter).Updates(request).Error
	if err != nil {
		cusErr := errors.DbError(err)
		return cusErr
	}
	return errors.CustomError{}
}

func (repo *Repository) CountryCallingCodeByConditionRepository(filter map[string]interface{}) (res []*models.CountriesCallingCode, cusErr errors.CustomError) {
	err := repo.repository.Model(&models.CountriesCallingCode{}).Where(filter).Find(&res).Error
	if err != nil {
		cusErr = errors.DbError(err)
		return nil, cusErr
	}
	return res, cusErr
}

func (repo *Repository) CountryCallingCodeByCharacterRepository(countryNameLike string) (res []*models.CountriesCallingCode, cusErr errors.CustomError) {
	err := repo.repository.Model(&models.CountriesCallingCode{}).Where("country_name LIKE ? ", "%"+countryNameLike+"%").Find(&res).Error
	if err != nil {
		cusErr = errors.DbError(err)
		return nil, cusErr
	}
	return res, cusErr
}

func SearchProducts(filter map[string][]string) (res []models.Product) {
	var repo *Repository
	err := repo.repository.Model(&models.Product{}).Find(&res).Error
	fmt.Println(err)
	return res
}
