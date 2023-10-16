package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dsn string) error {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = database.AutoMigrate(&Book{}, &Author{})
	if err != nil {
		return err
	}
	DB = database
	return nil
}
