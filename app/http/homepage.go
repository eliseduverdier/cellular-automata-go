package http

// Not used for the moment

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// Page with data to display
type Page struct {
}

// RenderHomePage renders a page with checkboxes, that will determine the first line
func RenderHomePage(w http.ResponseWriter, req *http.Request) {
	page := &Page{}

	cwd, _ := os.Getwd()
	templatePath := filepath.Join(cwd, "app/renderer/templates/home.html")

	t, _ := template.ParseFiles(templatePath)
	err := t.Execute(w, page)
	if err != nil {
		panic(err)
	}
}
