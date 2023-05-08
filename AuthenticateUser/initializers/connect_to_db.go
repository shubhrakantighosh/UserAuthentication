package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func configDSN() string {
	host := os.Getenv("host")
	user := os.Getenv("username")
	password := os.Getenv("password")
	dbname := os.Getenv("databaseName")
	port := os.Getenv("port")
	sslmode := os.Getenv("sslmode")

	return "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode
}

func ConnectToDB() *gorm.DB {
	dsn := configDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db.")
	}
	return db
}
