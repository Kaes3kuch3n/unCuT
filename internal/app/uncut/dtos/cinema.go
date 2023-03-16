package dtos

import "uncut/internal/app/uncut/entities"

type Cinema struct {
	Name       string      `json:"name"`
	Screenings []Screening `json:"screenings"`
}

func CinemaToDTO(c *entities.Cinema) (cinema Cinema) {
	screenings := make([]Screening, 0, len(c.Screenings))
	for _, screening := range c.Screenings {
		screenings = append(screenings, ScreeningToDTO(screening))
	}

	return Cinema{
		Name:       c.Name,
		Screenings: screenings,
	}
}
