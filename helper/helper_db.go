package helper

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(dsn string) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database, please check your database or connection string!")
	}
	fmt.Println("Successfully connected to database!")
	return DB
}
