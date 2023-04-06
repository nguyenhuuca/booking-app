package main

import (
	"booking-service/api"
	"booking-service/config"
)

func main() {
	globalConfig := config.LoadConfig()
	api.Router().Run("localhost:" + globalConfig.Port)
}
