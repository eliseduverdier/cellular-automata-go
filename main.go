package main

import (
	"flag"

	"github.com/eliseduverdier/cellular-automata-go/app/writer/writer_http"
	"github.com/eliseduverdier/cellular-automata-go/app/writer/writer_shell"
)

func main() {
	displayInShell := flag.Bool("shell", false, "If present, create the image in images/, or display the textual automata directly in the shell")
	flag.Parse()

	if *displayInShell {
		writer_shell.RenderShell()
	} else {
		writer_http.RenderHttp()
	}
}
