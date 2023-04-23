package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"uncut/internal/app/uncut/db"
)

func main() {
	database, err := gorm.Open(sqlite.Open("db.sqlite"))
	if err != nil {
		panic("failed to connect to database")
	}

	err = database.AutoMigrate(
		&db.Advertiser{},
		&db.Advertisement{},
		&db.Cinema{},
		&db.Movie{},
		&db.MovieScreen{},
		&db.Screen{},
		&db.Screening{},
		&db.ScreenType{},
	)
	if err != nil {
		panic("failed to migrate the database")
	}
}
