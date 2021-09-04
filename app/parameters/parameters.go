package parameters

import (
	"flag"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/eliseduverdier/cellular-automata-go/app/automata"
)

type Parameters struct {
	States      int
	Order       int
	Columns     int
	Rows        int
	RandomStart bool
	Rule        int
	Render      string
}

func GetFromRequest(req *http.Request) Parameters {
	states, _ := strconv.Atoi(req.URL.Query().Get("s"))
	order, _ := strconv.Atoi(req.URL.Query().Get("o"))
	width, _ := strconv.Atoi(req.URL.Query().Get("w"))
	height, _ := strconv.Atoi(req.URL.Query().Get("h"))
	rule, _ := strconv.Atoi(req.URL.Query().Get("r"))
	randomStart, _ := strconv.ParseBool(req.URL.Query().Get("random_start"))
	// Set defaults, TODO save elsewhere
	if states == 0 {
		states = 2
	}
	if order == 0 {
		order = 1
	}
	if width == 0 {
		width = 100
	}
	if height == 0 {
		height = 100
	}
	if rule == 0 {
		rule = rand.Intn(automata.GetMaxRule(automata.GetMaxStates(states, order)))
	}

	return Parameters{
		States:      states,
		Order:       order,
		Columns:     width,
		Rows:        height,
		Rule:        rule,
		RandomStart: randomStart,
		Render:      "image",
	}
}

func GetFromShell() Parameters {
	states := flag.Int("s", 2, "the number of states")
	order := flag.Int("o", 1, "the order (1 or 2)")
	columns := flag.Int("w", 100, "the number of columns")
	rows := flag.Int("h", 100, "the number of rows")
	rule := flag.Int("r", 73, "the rule number")
	randomStart := flag.Bool("random", true, "if the first line is random")
	render := flag.String("render", "image", "image|text")

	flag.Parse()

	return Parameters{
		States:      *states,
		Order:       *order,
		Columns:     *columns,
		Rows:        *rows,
		RandomStart: *randomStart,
		Rule:        *rule,
		Render:      *render,
	}
}