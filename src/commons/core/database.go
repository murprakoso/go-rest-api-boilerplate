package core

import (
	"fmt"
	commons "go-rest-api-boilerplate/src/commons/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitializeDatabase() {
	//dsn := "root:@tcp(127.0.0.1:3306)/go_rest_api?charset=utf8mb4&parseTime=True&loc=Local"
	config := commons.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Database connected")
	// db.AutoMigrate(&entity.Product{})

	DB = db
}
