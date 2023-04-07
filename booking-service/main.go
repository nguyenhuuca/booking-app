package main

import (
	"booking-service/api"
	"booking-service/config"
)

func main() {
	globalConfig := config.LoadConfig()
	err := api.Router().Run("localhost:" + globalConfig.Port)
	if err != nil {
		return
	}
}
