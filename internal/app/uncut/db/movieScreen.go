package db

import (
	"gorm.io/gorm"
	"uncut/internal/app/uncut/entities"
)

type MovieScreen struct {
	ID       uint `gorm:"primaryKey"`
	CinemaID uint
	Cinema   Cinema
	MovieID  uint
	Movie    Movie
	Path     string
}

func (m *MovieScreen) toEntity() (movieScreen *entities.MovieScreen) {
	return &entities.MovieScreen{FilePath: m.Path}
}

func GetScreensForMovies(db *gorm.DB, cinema *entities.Cinema, movies []*entities.Movie) (movieScreens []*entities.MovieScreen) {
	movieScreens = make([]*entities.MovieScreen, len(movies))
	for i, movie := range movies {
		movieScreens[i] = getScreenForMovie(db, cinema, movie)
	}
	return movieScreens
}

func getScreenForMovie(db *gorm.DB, cinema *entities.Cinema, movie *entities.Movie) (movieScreen *entities.MovieScreen) {
	var screen *MovieScreen
	db.Where("cinema_id = ? AND movie_id = ?", cinema.ID, movie.ID).Find(&screen)
	return screen.toEntity()
}
