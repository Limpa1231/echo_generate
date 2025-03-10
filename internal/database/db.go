package database

import (
	"firstRest/orm"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=youruser dbname=yourdb sslmode=disable password=yourpassword"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Автомиграция схемы
	err = DB.AutoMigrate(&orm.Message{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
