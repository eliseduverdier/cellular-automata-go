package writer_http

import (
	"image/png"
	"net/http"

	"github.com/eliseduverdier/cellular-automata-go/app"
	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
)

// RenderImagePage renders a PNG image in the request
func RenderImagePage(w http.ResponseWriter, req *http.Request) {
	image := app.RenderImage(parameters.GetFromRequest(req))

	w.Header().Set("Content-Type", "image/png")
	err := png.Encode(w, image)
	if err != nil {
		panic("Image couldn’t be encoded")
	}
}
