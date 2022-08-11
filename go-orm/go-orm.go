package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Product struct{
	gorm.Model
	Code string
	Name string
	Price uint
}

func main(){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Faild to connect database")
	}

	// db.AutoMigrate(&Product{})

	// db.Create(&Product{Code: "0001",Price: 200, Name: "洗发水"})
	db.Find
}