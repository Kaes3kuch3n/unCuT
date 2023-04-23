package lead

import (
	"errors"
	"fmt"
	"os"
	"uncut/internal/app/ffmpeg"
	"uncut/internal/app/uncut/entities"
)

type ResourceType int

const (
	Video ResourceType = iota
	Image
)

const defaultImageDuration = 5

type Resource struct {
	filePath     string
	resourceType ResourceType
	duration     float64
}

func GetTrailerResources(m *entities.Movie, movieScreen *entities.MovieScreen, upcomingScreen *entities.Screen) (resources []*Resource, err error) {
	var r *Resource
	r, err = NewImageResource(upcomingScreen.FilePath, defaultImageDuration)
	if err != nil {
		return nil, err
	}
	resources = append(resources, r)
	r, err = NewVideoResource(m.TrailerPath)
	if err != nil {
		return nil, err
	}
	resources = append(resources, r)
	r, err = NewImageResource(movieScreen.FilePath, defaultImageDuration)
	if err != nil {
		return nil, err
	}
	resources = append(resources, r)
	return resources, nil
}

func AdToResource(ad entities.Ad) (r *Resource, err error) {
	return AdToResourceWithDuration(ad, defaultImageDuration)
}

func AdToResourceWithDuration(ad entities.Ad, duration float64) (r *Resource, err error) {
	if ad.Type == entities.Image {
		r, err = NewImageResource(ad.FilePath, duration)
	} else if ad.Type == entities.Video {
		r, err = NewVideoResource(ad.FilePath)
	} else {
		panic(fmt.Errorf("unknown advertisement type: %s", ad.Type))
	}
	if err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func NewVideoResource(filePath string) (r *Resource, err error) {
	if !isValidFilePath(filePath) {
		return nil, errors.New(fmt.Sprintf("file does not exist: %s", filePath))
	}
	duration, err := getResourceDuration(filePath)
	if err != nil {
		return nil, err
	}
	return &Resource{filePath: filePath, resourceType: Video, duration: duration}, nil
}

func NewImageResource(filePath string, duration float64) (r *Resource, err error) {
	if !isValidFilePath(filePath) {
		return nil, errors.New(fmt.Sprintf("file does not exist: %s", filePath))
	}
	return &Resource{filePath: filePath, resourceType: Image, duration: duration}, nil
}

func (r *Resource) GetWeight() float64 {
	return r.duration
}

func (r *Resource) ToImport() *ffmpeg.Import {
	switch r.resourceType {
	case Image:
		return ffmpeg.NewImageImport(r.filePath, r.duration)
	case Video:
		return ffmpeg.NewVideoImport(r.filePath)
	default:
		panic("unknown resource type")
	}
}

func isValidFilePath(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func getResourceDuration(filePath string) (duration float64, err error) {
	d, err := ffmpeg.GetDuration(filePath)
	if err != nil {
		return -1, err
	}
	return d, nil
}
