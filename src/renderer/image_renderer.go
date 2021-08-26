package renderer

import (
	"os"
	"image"
	"image/color"
	"image/png"
)

func GenerateImage(matrix [][]int) {

	width := len(matrix[0])
	height := len(matrix)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	black := color.RGBA{0, 0, 0, 0xff}


	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case matrix[x][y] == 0:
				img.Set(y, x, black)
			case matrix[x][y] == 1:
				img.Set(y, x, color.White)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}