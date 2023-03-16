package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(dsn string) (database *gorm.DB) {
	database, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("failed to connect to database [%w]", err))
	}
	return database
}
