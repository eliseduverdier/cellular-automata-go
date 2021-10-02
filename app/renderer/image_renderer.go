package renderer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// GenerateImage  saves a PNG from the cellular automata's matrix
func GenerateImage(matrix [][]int, metadata map[string]int, destination string) (*image.RGBA, string) {

	width := len(matrix[0])
	height := len(matrix)

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

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			img.Set(y, x, colors[matrix[x][y]])
		}
	}

	filename := GetImageName(metadata)
	f, err := os.Create(destination + "/" + filename + ".png")
	if err != nil {
		panic(fmt.Sprintf("Could not create image file: %v", err))
	}

	err = png.Encode(f, img)
	if err != nil {
		panic("Could not generate image")
	}

	return img, filename
}

func GetImageName(metadata map[string]int) string {
	return fmt.Sprintf("s%d-o%d-r%d", metadata["states"], metadata["order"], metadata["rule"])
}
