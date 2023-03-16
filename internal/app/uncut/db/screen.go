package db

import (
	"uncut/internal/app/uncut/entities"
)

type Screen struct {
	ID           uint `gorm:"primaryKey"`
	CinemaID     uint
	Cinema       Cinema
	ScreenTypeID uint
	ScreenType   ScreenType
	Path         string
}

func (s *Screen) toEntity() (screen *entities.Screen) {
	return &entities.Screen{FilePath: s.Path}
}
