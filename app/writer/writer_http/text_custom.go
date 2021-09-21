package writer_http

// Not used for the moment

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// Page with data to display
type Page struct {
	Checkboxes []int
}

// RenderCustomTextPage renders a page with checkboxes, that will determine the first line
func RenderCustomTextPage(w http.ResponseWriter, req *http.Request) {
	// TODO get first line array from req.URL.Query().Get("i")

	checkboxes := make([]int, 50)
	for i := range checkboxes {
		checkboxes[i] = i
	}

	page := &Page{checkboxes}

	cwd, _ := os.Getwd()
	templatePath := filepath.Join(cwd, "app/renderer/templates/custom.html")

	t, _ := template.ParseFiles(templatePath)
	err := t.Execute(w, page)
	if err != nil {
		panic(err) // ...
	}
}
