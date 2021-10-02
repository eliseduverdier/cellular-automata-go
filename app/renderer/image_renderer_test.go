package renderer

import (
	_ "fmt"
	_ "os"
	"testing"

	_ "github.com/eliseduverdier/cellular-automata-go/app/automata"
)

func TestImage(t *testing.T) {
	// todo image testdata: s2, c1, w5, h5, r110, FirstLine00100
	// 	metadata := map[string]int{
	// 		"states": 2,
	// 		"order":  1,
	// 		"rule":   110,
	// 	}
	// 	actual := GenerateImage([][]int{
	// 		{0, 0, 1, 0, 0},
	// 		{0, 1, 1, 0, 0},
	// 		{1, 1, 1, 0, 0},
	// 		{1, 0, 1, 0, 1},
	// 		{1, 1, 1, 1, 1},
	// 	}, metadata)

	// 	expectedPath := "/../tests/testdata/s2-o1-r110-w5.png"
	// 	imgFile, err := os.Open(expectedPath)

	// 	if err != nil {
	// 		t.Errorf("Image '%s' was not found", expectedPath)
	// 	}

	// 	img, filename, _ := image.Decode(io.Rea(expectedPath))
}
