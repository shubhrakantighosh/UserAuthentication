package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	err := godotenv.Load("config/config.env")
	if err != nil {
		log.Fatalf("Failed to open config.env file. %v", err)
	}
}
