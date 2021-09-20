package writer_http

import (
	"fmt"
	"net/http"
)

func RenderHttp() {
	fmt.Println(" ->> Go to http://localhost:8888/text ",
		"or http://localhost:8888/image and tweak parameters",
		"or http://localhost:8888/custom and create your own")
	http.HandleFunc("/text", RenderTextPage)
	http.HandleFunc("/custom", RenderCustomTextPage)
	http.HandleFunc("/image", RenderImagePage)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic("Server doesn't run")
	}
}
