package database

import (
	"fmt"
	models "tugaswp/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Connect to the database
	sql, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/tugaswpbe?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		fmt.Println("Not Connected")
	}

	DB = sql

	DB.AutoMigrate(models.Product{})
	DB.AutoMigrate(models.User{})
	DB.AutoMigrate(models.ProductHistory{})
}
