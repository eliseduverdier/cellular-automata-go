package renderer

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// GenerateImage: saves a PNG from the cellular automata's matrix
func GenerateImage(matrix [][]int, imageName string) {

	width := len(matrix)
	height := len(matrix[0])

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

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

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, colors[matrix[x][y]])
		}
	}

	// Encode as PNG.
	f, _ := os.Create("images/" + imageName + ".png")
	png.Encode(f, img)
}
