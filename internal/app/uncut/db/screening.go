package db

import (
	"gorm.io/gorm"
	"time"
	"uncut/internal/app/uncut/entities"
)

type Screening struct {
	ID       uint `gorm:"primaryKey"`
	CinemaID uint
	Cinema   Cinema
	MovieID  uint
	Movie    Movie
	Date     time.Time
}

func (s *Screening) toEntity() (screening *entities.Screening) {
	return &entities.Screening{
		ID:     s.ID,
		Cinema: s.Cinema.toEntity(),
		Movie:  s.Movie.toEntity(),
		Date:   s.Date,
	}
}

func LoadScreening(db *gorm.DB, id uint) (screening *entities.Screening) {
	var s Screening
	db.Preload("Cinema.Screens.ScreenType").Preload("Movie").Find(&s, id)
	return s.toEntity()
}
