package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.Logger.LogMode(1)
	DB.AutoMigrate(Product{})

	DB.Create(&Product{})
	a := Product{}
	DB.First(&a)
}
