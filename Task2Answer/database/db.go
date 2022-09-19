package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var Instance DBInstance

func Connect() {
	dsn := "host=localhost user=hya password=knightfall dbname=spotlas_backend_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("[Error]: Error while connecting to the database.")
		os.Exit(1)
	}

	log.Println("[Success]: Successfully Connected to the database.")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("[Info]: Running Migrations.")
	db.AutoMigrate()

	Instance = DBInstance{
		Db: db,
	}

}
