package main

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"
)

func connectDB() *gorm.DB {
	db, err := gorm.Open(sqlserver.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		fmt.Println("Orrrrrr Nooooo")
	}
	return db
}
