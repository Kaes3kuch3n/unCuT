package dtos

import (
	"time"
	"uncut/internal/app/uncut/entities"
)

type Screening struct {
	ID    uint      `json:"id"`
	Movie string    `json:"movie"`
	Date  time.Time `json:"date"`
}

func ScreeningToDTO(s *entities.Screening) (screening Screening) {
	return Screening{
		ID:    s.ID,
		Movie: s.Movie.Name,
		Date:  s.Date,
	}
}
