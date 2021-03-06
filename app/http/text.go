package http

import (
	"fmt"
	"net/http"

	"github.com/eliseduverdier/cellular-automata-go/app"
	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
)

// RenderTextPage Displays the automata as a text in an html page
func RenderTextPage(w http.ResponseWriter, req *http.Request) {
	image := app.RenderText(parameters.GetFromRequest(req))

	req.Header.Add("Application", "text/html")
	fmt.Fprintf(w,
		"<html><body><pre>%s</pre></body></html>",
		image,
	)
}
