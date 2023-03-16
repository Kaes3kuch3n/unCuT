package entities

type AdType string

// AdMap maps the ad ID to its corresponding advertisement
type AdMap map[uint]Ad

const (
	Image AdType = "IMAGE"
	Video AdType = "VIDEO"
)

type Ad struct {
	ID         uint
	Name       string
	Type       AdType
	FilePath   string
	Advertiser *Advertiser
}
