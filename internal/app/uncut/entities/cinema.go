package entities

type ScreenMap map[uint]*Screen

type Cinema struct {
	ID         uint
	Name       string
	Screens    ScreenMap
	Screenings []*Screening
}
