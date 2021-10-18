package writer_http

import (
	"net/http"
	"os"
)

// RenderHttp Handles 3 routes to display by text or image
func RenderHttp() {
	// fmt.Println(" ->> Go to http://localhost:8888/text ",
	// 	"\n or http://localhost:8888/image and tweak parameters",
	// 	"\n or http://localhost:8888/custom and create your own")
	http.HandleFunc("/text", RenderTextPage)
	http.HandleFunc("/custom", RenderCustomTextPage)
	http.HandleFunc("/image", RenderImagePage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("Server doesn't run")
	}
}
