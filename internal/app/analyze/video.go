package analyze

import (
	"uncut/internal/app/ffmpeg"
)

type Video interface {
	FilePath() string
	Duration() float64
	FrameRate() float64
	FrameLength() float64
}

func NewVideo(filePath string) (v Video, err error) {
	duration, err := ffmpeg.GetDuration(filePath)
	fps, err := ffmpeg.GetFrameRate(filePath)
	if err != nil {
		return nil, err
	}

	return &video{
		filePath:  filePath,
		duration:  duration,
		frameRate: fps,
	}, nil
}

type video struct {
	filePath  string
	duration  float64
	frameRate float64
}

func (v *video) FilePath() string {
	return v.filePath
}

func (v *video) Duration() float64 {
	return v.duration
}

func (v *video) FrameRate() float64 {
	return v.frameRate
}

func (v *video) FrameLength() float64 {
	return 1.0 / v.frameRate
}
