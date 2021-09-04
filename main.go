package main

import (
	"flag"
	"fmt"
	"image/png"
	"net/http"

	"github.com/eliseduverdier/cellular-automata-go/app"
	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
)

func renderImagePage(w http.ResponseWriter, req *http.Request) {
	image := app.RenderImage(parameters.GetFromRequest(req))

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, image)
}

func renderTextPage(w http.ResponseWriter, req *http.Request) {
	image := app.RenderText(parameters.GetFromRequest(req))

	req.Header.Add("Application", "text/html")
	fmt.Fprintf(w,
		"<h1>Hello Cellular Automata</h1>"+
			"<pre>%s</pre>",
		image,
	)
}

func renderShell() {
	params := parameters.GetFromShell()
	app.Render(params)
}

func main() {
	// TODO doesn't work (:
	displayInShell := flag.Bool("shell", false, "If present, create the image in images/, or display the textual automata directly in the shell")

	if *displayInShell {
		renderShell()
	} else {
		fmt.Println(" ->> Go to http://localhost:8888/text or http://localhost:8888/image and tweak parameters")
		http.HandleFunc("/text", renderTextPage)
		http.HandleFunc("/image", renderImagePage)
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			panic("Server doesn't run")
		}
	}
}
