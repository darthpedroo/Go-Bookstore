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
	d, err := gorm.Open("mysql", "darthpedro:Axlesharma@12@/simplerest?") //Cambiar lo que dice Axel nose que 
	if err != nil {
		
	}
}