package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect database:", err)
	}

	fmt.Println("database connected")

	return db
}