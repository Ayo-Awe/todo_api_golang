package database

import (
	"log"
	"os"
	"todo_api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDB() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
		os.Exit(2)
	}

	log.Println("Connected to database")
	db.AutoMigrate(&models.Todo{})
	DBConn = db
}
