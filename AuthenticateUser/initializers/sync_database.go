package initializers

import (
	"AuthenticateUser/domain/models"
	"gorm.io/gorm"
	"log"
)

func SyncDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&models.Users{}, &models.UserSessions{})
	if err != nil {
		log.Fatalf("Faild to Migration %v", err)
	}
}
