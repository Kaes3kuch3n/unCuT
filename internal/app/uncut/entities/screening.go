package entities

import "time"

type Screening struct {
	ID     uint
	Cinema *Cinema
	Movie  *Movie
	Date   time.Time
}
