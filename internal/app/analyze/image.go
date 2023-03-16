package analyze

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func LoadImage(filePath string) (img image.Image, err error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	img, _, err = image.Decode(bytes.NewBuffer(file))
	if err != nil {
		return nil, err
	}
	return img, nil
}
