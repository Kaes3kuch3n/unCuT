package db

import "uncut/internal/app/uncut/entities"

type Advertiser struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Advertisements []Advertisement
}

func (a Advertiser) toEntity() (advertiser *entities.Advertiser) {
	return &entities.Advertiser{Name: a.Name}
}
