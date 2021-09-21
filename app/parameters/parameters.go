package parameters

import (
	"flag"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/eliseduverdier/cellular-automata-go/app/automata"
)

type Parameters struct {
	States  int
	Order   int
	Columns int
	Rows    int
	Start   []int
	Rule    int
	Render  string
}

func GetFromRequest(req *http.Request) Parameters {
	states, _ := strconv.Atoi(req.URL.Query().Get("s"))
	order, _ := strconv.Atoi(req.URL.Query().Get("o"))
	width, _ := strconv.Atoi(req.URL.Query().Get("w"))
	height, _ := strconv.Atoi(req.URL.Query().Get("h"))
	rule, _ := strconv.Atoi(req.URL.Query().Get("r"))
	firstLineType := req.URL.Query().Get("start") // will be used later to generate the first line early
	firstLineContent := req.URL.Query().Get("line")

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

	// Generate first line from appropriate parameters

	firstLine := make([]int, width)
	if firstLineType == "centered" {
		firstLine = automata.GetCenteredLine(width, states)
	} else if firstLineContent != "" {
		line := strings.Split(firstLineContent, "")
		for i, v := range line {
			firstLine[i], _ = strconv.Atoi(v)
		}
	} else {
		firstLine = automata.GetRandomLine(width, states)
	}

	return Parameters{
		States:  states,
		Order:   order,
		Columns: width,
		Rows:    height,
		Rule:    rule,
		Start:   firstLine,
		Render:  "image",
	}
}

func GetFromShell() Parameters {
	states := flag.Int("s", 2, "the number of states")
	order := flag.Int("o", 1, "the order (1 or 2)")
	width := flag.Int("w", 100, "the number of columns")
	height := flag.Int("h", 100, "the number of rows")
	rule := flag.Int("r", 73, "the rule number")

	firstLineType := flag.String("start", "random", "the type of first line if automatically generated (random, centered)")
	firstLineContent := flag.String("line", "1", "the content of the first line")

	firstLine := make([]int, *width)
	if *firstLineType == "centered" {
		firstLine = automata.GetCenteredLine(*width, *states)
	} else if *firstLineType == "random" {
		firstLine = automata.GetRandomLine(*width, *states)
	} else {
		line := strings.Split(*firstLineContent, "")
		for i, v := range line {
			firstLine[i], _ = strconv.Atoi(v)
		}
	}

	render := flag.String("render", "image", "image|text")

	flag.Parse() // FIXME: shell params were already parsed in shell.go>RenderShell, cannot parse again

	return Parameters{
		States:  *states,
		Order:   *order,
		Columns: *width,
		Rows:    *height,
		Start:   firstLine,
		Rule:    *rule,
		Render:  *render,
	}
}
