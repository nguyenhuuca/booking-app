package db

import (
	"booking-service/config"
	"booking-service/storage"
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

	log.Println("migration data...")
	err = db.AutoMigrate(&storage.Branch{}, &storage.Product{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return db
}
