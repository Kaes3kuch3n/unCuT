package db

import (
	"fmt"
	"gorm.io/gorm"
	"uncut/internal/app/uncut/entities"
)

type Cinema struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Screens    []Screen
	Screenings []Screening
}

// Converts the database cinema model into its corresponding entity
func (c *Cinema) toEntity() (cinema *entities.Cinema) {
	screenings := make([]*entities.Screening, 0, len(c.Screenings))
	for _, screening := range c.Screenings {
		screenings = append(screenings, screening.toEntity())
	}

	return &entities.Cinema{
		Name:       c.Name,
		Screens:    c.mapScreens(),
		Screenings: screenings,
	}
}

// Maps the screen types to the cinema's screen files
func (c *Cinema) mapScreens() (screens entities.ScreenMap) {
	screens = make(entities.ScreenMap, len(c.Screens))
	for _, s := range c.Screens {
		screens[s.ScreenType.ID] = s.toEntity()
	}
	return screens
}

func GetCinemasByName(db *gorm.DB, likeName string, maxCount int) (cinemas []*entities.Cinema) {
	var c []Cinema
	db.
		Preload("Screenings").
		Preload("Screenings.Movie").
		Where("cinemas.name LIKE ?", fmt.Sprintf("%%%s%%", likeName)).
		Limit(maxCount).
		Find(&c)

	cinemas = make([]*entities.Cinema, 0, len(c))
	for _, cinema := range c {
		cinemas = append(cinemas, cinema.toEntity())
	}
	return cinemas
}
