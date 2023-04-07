package main

import (
	"booking-service/api"
	"booking-service/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		log.Fatal("Error loading .env file")
	}

	globalConfig := config.LoadConfig()
	err := api.Router().Run("localhost:" + globalConfig.Port)
	if err != nil {
		return
	}
}
