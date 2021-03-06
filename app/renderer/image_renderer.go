package renderer

import (
	"fmt"
	"image"
	"image/color"
)

// GenerateImage saves a PNG from the cellular automata's matrix
func GenerateImage(matrix [][]int, pixelSize int, metadata map[string]int, destination string) (*image.RGBA, string) {

	width := len(matrix[0])
	height := len(matrix)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width * pixelSize, height * pixelSize}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	colors := []color.RGBA{
		{0, 0, 0, 0xff},
		{255, 255, 255, 0xff},
		{134, 235, 2, 0xff},
		{45, 144, 168, 0xff},
		{175, 109, 232, 0xff},
		{230, 0, 103, 0xff},
		{184, 77, 0, 0xff},
		{222, 222, 18, 0xff},
		{93, 150, 0, 0xff},
		{58, 252, 129, 0xff},
	}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			if pixelSize > 1 {
				drawRect(*img, y*pixelSize, x*pixelSize, pixelSize, colors[matrix[x][y]])
			} else {
				img.Set(y, x, colors[matrix[x][y]])
			}
		}
	}

	return img, GetImageName(metadata)
}

// GetImageName from number of states, order, and rule number
func GetImageName(metadata map[string]int) string {
	return fmt.Sprintf("s%d-o%d-r%d", metadata["states"], metadata["order"], metadata["rule"])
}

func drawRect(image image.RGBA, x, y, pixelSize int, color color.RGBA) {
	for i := x; i < x+pixelSize; i++ {
		for j := y; j < y+pixelSize; j++ {

			image.Set(i, j, color)
		}
	}
}
