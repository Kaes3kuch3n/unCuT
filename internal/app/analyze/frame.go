package analyze

import (
	"image"
	_ "image/jpeg"
	"math"
	"os/exec"
	"strconv"
	"sync"
)

func GetAvgColorDifference(v Video, stepSize float64) (avgDiff float64) {
	numFrames := int(math.Floor(v.Duration() / stepSize))
	colors := getAvgFrameColors(v, 0.0, numFrames, stepSize)

	totalDiff := 0.0
	for i := 1; i < numFrames; i++ {
		totalDiff += getColorDifference(colors[i-1], colors[i])
	}

	return totalDiff / float64(numFrames)
}

func GetAvgFrameDifference(v Video, stepSize float64) (flashiness float64) {
	numFrames := int(math.Floor(v.Duration() / stepSize))

	frames := make([]image.Image, numFrames)
	var wg sync.WaitGroup
	wg.Add(numFrames)
	for i := 0; i < numFrames; i++ {
		go func(i int) {
			defer wg.Done()
			frames[i] = extractFrame(v, float64(i)*stepSize)
		}(i)
	}
	wg.Wait()

	totalDiff := 0.0
	for i := 1; i < numFrames; i++ {
		totalDiff += getFrameDifference(frames[i-1], frames[i])
	}

	return totalDiff / float64(numFrames)
}

func GetAvgColorOverFrames(v Video, position float64, duration float64, stepSize float64) (avgColor RGBColor) {
	if position+duration > v.Duration() {
		panic("duration crosses video end")
	}

	numFrames := int(math.Floor(duration * v.FrameRate()))
	colors := getAvgFrameColors(v, position, numFrames, stepSize)

	return getAverageColor(colors)
}

func GetAvgImageColor(i image.Image) (avgColor RGBColor) {
	return getAverageImgColor(i)
}

func getAvgFrameColors(v Video, position float64, numFrames int, stepSize float64) (colors []RGBColor) {
	colors = make([]RGBColor, numFrames)

	var wg sync.WaitGroup
	wg.Add(numFrames)

	for i := 0; i < numFrames; i++ {
		go func(i int) {
			defer wg.Done()
			c := getAvgFrameColor(v, position+float64(i)*stepSize)
			colors[i] = c
		}(i)
	}

	wg.Wait()

	return colors
}

func getAvgFrameColor(v Video, position float64) (color RGBColor) {
	frame := extractFrame(v, position)
	return getAverageImgColor(frame)
}

func getFrameDifference(a image.Image, b image.Image) (diff float64) {
	bounds := a.Bounds()
	if bounds.Dx() != b.Bounds().Dx() || bounds.Dy() != b.Bounds().Dy() {
		panic("images passed to getFrameDifference must have equal bounds")
	}

	imgA := imageToRGBA(a)
	imgB := imageToRGBA(b)

	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			pos := y*bounds.Dy() + x
			aPixColor := imgA.Pix[pos*4 : pos*4+4]
			bPixColor := imgB.Pix[pos*4 : pos*4+4]
			diff += getColorDifference(ToRGBColor(aPixColor), ToRGBColor(bPixColor))
		}
	}

	return diff / float64(bounds.Dx()*bounds.Dy())
}

func extractFrame(v Video, position float64) (img image.Image) {
	// TODO: Figure out how to input file from stdin
	cmd := exec.Command("ffmpeg",
		"-ss", strconv.FormatFloat(position, 'f', -1, 64),
		"-i", v.FilePath(),
		"-frames:v", "1",
		"-update", "true",
		"-c:v", "mjpeg",
		"-f", "image2pipe",
		"-")

	imgPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	img, _, err = image.Decode(imgPipe)
	if err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}

	return img
}
