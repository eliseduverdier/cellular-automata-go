package renderer

import (
	"os"
	"image"
	"image/color"
	"image/png"
)

func GenerateImage(matrix [][]int, imageName string) {

	width := len(matrix[0])
	height := len(matrix)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	colors := []color.RGBA{
		color.RGBA{  0,   0,   0, 0xff},
		color.RGBA{255, 255, 255, 0xff},
		color.RGBA{134, 235,   2, 0xff},
		color.RGBA{ 45, 144, 168, 0xff},
		color.RGBA{175, 109, 232, 0xff},
		color.RGBA{230,   0, 103, 0xff},
		color.RGBA{184,  77,   0, 0xff},
		color.RGBA{222, 222,  18, 0xff},
		color.RGBA{ 93, 150,   0, 0xff},
		color.RGBA{ 58, 252, 129, 0xff},
	}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(y, x, colors[matrix[x][y]])
		}
	}

	// Encode as PNG.
	f, _ := os.Create("images/"+imageName+".png")
	png.Encode(f, img)
}