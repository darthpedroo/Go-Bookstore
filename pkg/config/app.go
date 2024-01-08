package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// This file purpose is to return a db variable which will interact with the other parts of the code.


//We use a pointer to change the db state
var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local") 
	
	if err != nil {
		panic(err)
	}

	db = d // NO RETURN BECAUSE THE VARIABLE IS A POINTER
}

func GetDB() *gorm.DB {
	return db
}