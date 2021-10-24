package renderer

import (
	_ "fmt"
	"image"
	"image/color"
	_ "os"
	"testing"

	_ "github.com/eliseduverdier/cellular-automata-go/app/automata"
)

func TestGenerateImage(t *testing.T) {
	actualImg, actualName := getImage()

	expectedImg := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{5, 5}})

	// compare name

	expectedName := "s2-o1-r110"
	if actualName != expectedName {
		t.Errorf("Image name is incorrect, expected %v, got %s.", expectedName, actualName)
	}

	// compare rectangle size
	if actualImg.Rect.Max != expectedImg.Rect.Max || actualImg.Rect.Min != expectedImg.Rect.Min {
		t.Errorf("Wrong image size, expected %v×%v, got %v×%v", 5, 5, actualImg.Rect.Min, actualImg.Rect.Max)
	}

	expectedColor := getPixelsColors()

	// compare each pixel color value
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if actualImg.RGBAAt(i, j) != expectedColor[j][i] {
				t.Errorf("Wrong 0,0 pixel color, expected %v, got %v", actualImg.RGBAAt(i, j), expectedColor[j][i])
			}
		}
	}

}

func getImage() (*image.RGBA, string) {
	metadata := map[string]int{
		"states": 2,
		"order":  1,
		"rule":   110,
	}
	return GenerateImage([][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
	}, metadata, "")
}

func getPixelsColors() [][]color.RGBA {
	b := color.RGBA{0, 0, 0, 0xff}
	w := color.RGBA{255, 255, 255, 0xff}

	return [][]color.RGBA{
		{b, b, w, b, b},
		{b, w, w, b, b},
		{w, w, w, b, b},
		{w, b, w, b, w},
		{w, w, w, w, w},
	}
}
