package dtos

import "uncut/internal/app/uncut/entities"

type Ad struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Advertiser string `json:"advertiser"`
}

func AdToDTO(a entities.Ad) (ad Ad) {
	return Ad{
		ID:         a.ID,
		Name:       a.Name,
		Type:       string(a.Type),
		Advertiser: a.Advertiser.Name,
	}
}
