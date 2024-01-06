package database

import (
  "fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/go_rest_api?charset=utf8mb4&parseTime=True&loc=Local"
  // db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Can't connect to database")
	}

	fmt.Println("Connected to Database")
}
