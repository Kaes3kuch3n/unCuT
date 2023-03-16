package db

import (
	"gorm.io/gorm"
	"uncut/internal/app/uncut/entities"
)

type ScreenType struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (s ScreenType) toEntity() (screen entities.ScreenType) {
	return entities.ScreenType{ID: s.ID, Name: s.Name}
}

func LoadScreenTypes(database *gorm.DB) (types []entities.ScreenType) {
	var s []ScreenType
	database.Find(&s)

	types = make([]entities.ScreenType, 0, len(s))
	for _, t := range s {
		types = append(types, t.toEntity())
	}
	return types
}
