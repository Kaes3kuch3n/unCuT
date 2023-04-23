package db

import (
	"gorm.io/gorm"
	"uncut/internal/app/uncut/entities"
)

type Movie struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	TrailerPath string
}

func (m *Movie) toEntity() (movie *entities.Movie) {
	return &entities.Movie{
		ID:          m.ID,
		Name:        m.Name,
		TrailerPath: m.TrailerPath,
	}
}

func GetUpcomingMovies(db *gorm.DB, screening *entities.Screening, count int) (movies []*entities.Movie) {
	results := make([]map[string]interface{}, 0, count)
	db.Model(&Screening{}).
		Joins("Movie").
		Limit(count).
		Order("date").
		Where("date > ?", screening.Date.UnixMilli()).
		Find(&results)

	movies = make([]*entities.Movie, 0, count)
	for _, movie := range results {
		movie := &Movie{
			ID:          uint(movie["Movie__id"].(int64)),
			Name:        movie["Movie__name"].(string),
			TrailerPath: movie["Movie__trailer_path"].(string),
		}
		movies = append(movies, movie.toEntity())
	}
	return movies
}
