package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/eliseduverdier/cellular-automata-go/app/writer"
)

func main() {
	displayInShell := flag.Bool("shell", false, "If present, create the image in images/, or display the textual automata directly in the shell")
	flag.Parse()

	if *displayInShell {
		writer.RenderShell()
	} else {
		fmt.Println(" ->> Go to http://localhost:8888/text or http://localhost:8888/image and tweak parameters")
		http.HandleFunc("/text", writer.RenderTextPage)
		http.HandleFunc("/image", writer.RenderImagePage)
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			panic("Server doesn't run")
		}
	}
}
