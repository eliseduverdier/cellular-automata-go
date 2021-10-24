package http

import (
	"fmt"
	"net/http"
	"os"
)

// RenderHttp Handles routes
func RenderHttp() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fmt.Println(" ->>  http://localhost:"+port+"/",
		"\n ->> http://localhost:"+port+"/image ",
		"\n ->>  http://localhost:"+port+"/text")

	http.HandleFunc("/", RenderHomePage)
	http.HandleFunc("/text", RenderTextPage)
	http.HandleFunc("/image", RenderImagePage)
	http.HandleFunc("/favicon.ico", RenderFavicon)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
