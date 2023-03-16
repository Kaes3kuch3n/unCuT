package db

import (
	"gorm.io/gorm"
	"uncut/internal/app/uncut/entities"
)

type Advertisement struct {
	ID           uint `gorm:"primaryKey"`
	AdvertiserID uint
	Advertiser   Advertiser
	Name         string
	Type         string
	Path         string
}

func (a *Advertisement) toEntity() (ad entities.Ad) {
	return entities.Ad{
		ID:         a.ID,
		Name:       a.Name,
		Type:       entities.AdType(a.Type),
		FilePath:   a.Path,
		Advertiser: a.Advertiser.toEntity(),
	}
}

func LoadAds(db *gorm.DB) (ads entities.AdMap) {
	advertisements := make([]Advertisement, 10)
	db.Joins("Advertiser").Find(&advertisements)

	ads = make(entities.AdMap, len(advertisements))
	for _, ad := range advertisements {
		ads[ad.ID] = ad.toEntity()
	}
	return ads
}
