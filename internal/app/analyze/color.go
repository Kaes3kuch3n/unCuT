package analyze

import (
	"image"
	"image/draw"
	"math"
)

type RGBColor struct {
	R uint
	G uint
	B uint
}

func ToRGBColor(c []uint8) (color RGBColor) {
	return RGBColor{
		R: uint(c[0]),
		G: uint(c[1]),
		B: uint(c[2]),
	}
}

func GetPerceivedLightness(color RGBColor) (lightness float64) {
	vR := linearize(float64(color.R))
	vG := linearize(float64(color.G))
	vB := linearize(float64(color.B))

	Y := 0.2126*vR + 0.7152*vG + 0.0722*vB

	if Y <= (216.0 / 24389.0) {
		return Y * (24389.0 / 27.0)
	} else {
		return math.Pow(Y, 1.0/3.0)*116.0 - 16.0
	}
}

func getColorDifference(a RGBColor, b RGBColor) (diff float64) {
	diffR := math.Abs(float64(a.R) - float64(b.R))
	diffG := math.Abs(float64(a.G) - float64(b.G))
	diffB := math.Abs(float64(a.B) - float64(b.B))

	return (diffR + diffG + diffB) / 3.0
}

func getAverageImgColor(img image.Image) (avgColor RGBColor) {
	r := uint(0)
	g := uint(0)
	b := uint(0)

	bounds := img.Bounds()
	numPixels := uint(bounds.Dx() * bounds.Dy())
	rgba := imageToRGBA(img)

	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			pos := y*bounds.Dy() + x
			pix := rgba.Pix[pos*4 : pos*4+4]
			r += uint(pix[0])
			g += uint(pix[1])
			b += uint(pix[2])
		}
	}

	return RGBColor{
		R: r / numPixels,
		G: g / numPixels,
		B: b / numPixels,
	}
}

func imageToRGBA(i image.Image) (img *image.RGBA) {
	rgba := image.NewRGBA(i.Bounds())
	draw.Draw(rgba, i.Bounds(), i, image.Point{}, draw.Src)
	return rgba
}

func getAverageColor(colors []RGBColor) (avg RGBColor) {
	n := uint(len(colors))
	c := RGBColor{R: 0, G: 0, B: 0}

	for i := uint(0); i < n; i++ {
		c.R += colors[i].R
		c.G += colors[i].G
		c.B += colors[i].B
	}

	c.R /= n
	c.G /= n
	c.B /= n

	return c
}

func linearize(v float64) (vLin float64) {
	v = v / 255.0

	if v <= 0.04045 {
		return v / 12.92
	} else {
		return math.Pow((v+0.055)/1.055, 2.4)
	}
}
