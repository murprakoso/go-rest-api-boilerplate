package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitializeDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_rest_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Database connected")
	//db.AutoMigrate(&entity.Product{})

	DB = db
}
