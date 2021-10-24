package http

import (
	"image/png"
	"math/rand"
	"net/http"

	"github.com/eliseduverdier/cellular-automata-go/app"
	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
)

// Favicon generates favicon
func RenderFavicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "immutable")

	goodChoices := []int{73, 110, 150, 160}

	params := parameters.Parameters{
		States:  2,
		Order:   1,
		Columns: 16,
		Rows:    16,
		Start:   []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		Rule:    float64(goodChoices[rand.Intn(len(goodChoices))]),
		Render:  "image",
	}
	image, _ := app.RenderImage(params)

	err := png.Encode(w, image)
	if err != nil {
		panic("Image couldn’t be encoded")
	}
}
