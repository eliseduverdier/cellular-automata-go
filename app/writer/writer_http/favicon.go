package writer_http

import (
	"image/png"
	"net/http"

	"github.com/eliseduverdier/cellular-automata-go/app"
	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
)

// Favicon generates favicon
func Favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "immutable")

	params := parameters.Parameters{
		States:  2,
		Order:   1,
		Columns: 16,
		Rows:    16,
		Start:   []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		Rule:    150,
		Render:  "image",
	}
	image, name := app.RenderImage(params)

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Filename", name)
	err := png.Encode(w, image)
	if err != nil {
		panic("Image couldn’t be encoded")
	}
}
