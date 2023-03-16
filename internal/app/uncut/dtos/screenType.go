package dtos

import "uncut/internal/app/uncut/entities"

type ScreenType struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ScreenTypeToDTO(s entities.ScreenType) (screen ScreenType) {
	return ScreenType{ID: s.ID, Name: s.Name}
}
