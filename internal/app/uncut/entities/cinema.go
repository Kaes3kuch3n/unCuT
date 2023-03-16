package entities

type ScreenMap map[uint]*Screen

type Cinema struct {
	Name       string
	Screens    ScreenMap
	Screenings []*Screening
}
