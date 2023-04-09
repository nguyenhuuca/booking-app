package db

import (
	"audit-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	conf := config.LoadConfig()
	db, err := gorm.Open(postgres.Open(conf.DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connect db success...")

	return db
}
