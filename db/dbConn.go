package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDb() {
	var err error
	dsn := "host=localhost user=postgres password=adminPassword dbname=GO_USER_JWT port=5432 sslmode=disable"
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect db")
	}
	fmt.Println("Connected to DB")
}
