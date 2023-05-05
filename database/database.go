package database

import (
	"fmt"
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
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
		os.Exit(2)
	}

	log.Println("Connected to database")
	db.AutoMigrate(&models.Todo{})
	DBConn = db
}
