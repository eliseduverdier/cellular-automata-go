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
	metadata := map[string]int{
		"states": 2,
		"order":  1,
		"rule":   110,
	}
	actualImg, actualName := GenerateImage([][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
	}, 1, metadata, "")

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

	b := color.RGBA{0, 0, 0, 0xff}
	w := color.RGBA{255, 255, 255, 0xff}

	expectedColor := [][]color.RGBA{
		{b, b, w, b, b},
		{b, w, w, b, b},
		{w, w, w, b, b},
		{w, b, w, b, w},
		{w, w, w, w, w},
	}

	// compare each pixel color value
	compareTwoImagesPixelByPixels(actualImg, expectedColor, t)
}

func TestGetBiggerPixelSize(t *testing.T) {
	metadata := map[string]int{
		"states": 2,
		"order":  1,
		"rule":   110,
	}
	actual, _ := GenerateImage([][]int{
		{0, 1},
		{1, 0},
	}, 3, metadata, "")

	b := color.RGBA{0, 0, 0, 0xff}
	w := color.RGBA{255, 255, 255, 0xff}

	expected := [][]color.RGBA{
		{b, b, b, w, w, w},
		{b, b, b, w, w, w},
		{b, b, b, w, w, w},
		{w, w, w, b, b, b},
		{w, w, w, b, b, b},
		{w, w, w, b, b, b},
	}
	compareTwoImagesPixelByPixels(actual, expected, t)
}

func compareTwoImagesPixelByPixels(actual *image.RGBA, expected [][]color.RGBA, t *testing.T) {
	for i := 1; i < len(expected[0]); i++ {
		for j := 1; j < len(expected); j++ {
			if actual.RGBAAt(i, j) != expected[j][i] {
				t.Errorf("Wrong %d,%d pixel color, expected %v, got %v", i, j, actual.RGBAAt(i, j), expected[j][i])
			}
		}
	}
}
