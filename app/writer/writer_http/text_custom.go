package writer_http

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Page struct {
	Checkboxes []int
}

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
	t.Execute(w, page)
}
