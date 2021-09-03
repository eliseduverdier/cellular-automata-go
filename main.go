package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/eliseduverdier/cellular-automata-go/app"
	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "here automata \n")
}

func main() {
	params := getShellParameters()

	app.Render(params)

	// TODO getHttpGetParemeters()

	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic("Server doesnt run")
	}
}

func getShellParameters() parameters.Parameters {
	states := flag.Int("s", 2, "the number of states")
	order := flag.Int("o", 1, "the order (1 or 2)")
	columns := flag.Int("w", 100, "the number of columns")
	rows := flag.Int("h", 100, "the number of rows")
	rule := flag.Int("r", 73, "the rule number")
	randomStart := flag.Bool("random", true, "if the first line is random")
	render := flag.String("render", "image", "image|text")

	flag.Parse()

	return parameters.Parameters{
		States:      *states,
		Order:       *order,
		Columns:     *columns,
		Rows:        *rows,
		RandomStart: *randomStart,
		Rule:        *rule,
		Render:      *render,
	}
}
