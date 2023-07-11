package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"quote/internal/models"
)

type Conn struct {
	DB *gorm.DB
}

var Database Conn

func DBconn() *gorm.DB {
	url := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// TODO: Add Migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = Conn{DB: db}
	return db
}
