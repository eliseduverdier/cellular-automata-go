package parameters

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/eliseduverdier/cellular-automata-go/app/automata"
)

// Parameters to tweak the automata
type Parameters struct {
	States  int
	Order   int
	Columns int
	Rows    int
	Start   []int
	Rule    float64
	Render  string
}

// GetFromRequest Parse query string to get the parameters
func GetFromRequest(req *http.Request) Parameters {
	states, _ := strconv.Atoi(req.URL.Query().Get("s"))
	order, _ := strconv.Atoi(req.URL.Query().Get("o"))
	width, _ := strconv.Atoi(req.URL.Query().Get("w"))
	height, _ := strconv.Atoi(req.URL.Query().Get("h"))
	rule, _ := strconv.ParseInt(req.URL.Query().Get("r"), 10, 64)
	ruleNb := float64(rule)
	firstLineType := req.URL.Query().Get("start") // will be used later to generate the first line early
	firstLineContent := req.URL.Query().Get("line")
	renderType := req.URL.Path[1:]

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
		height = width
	}
	if rule == 0 {
		max := float64(automata.GetMaxRule(automata.GetMaxStatesCombinaisons(states, order)))
		ruleNb = rand.Float64() * max
	}

	// Generate first line from appropriate parameters

	firstLine := make([]int, width)
	switch firstLineType {
	case "centered":
		firstLine = automata.GetCenteredLine(width, states)
	case "right":
		firstLine = automata.GetPixelsOnRight(width, states)
	case "left":
		firstLine = automata.GetPixelsOnLeft(width, states)
	default:
		if firstLineContent != "" {
			line := strings.Split(firstLineContent, "")
			for i, v := range line {
				firstLine[i], _ = strconv.Atoi(v)
			}
		} else {
			firstLine = automata.GetRandomLine(width, states)
		}
	}

	return Parameters{
		States:  states,
		Order:   order,
		Columns: width,
		Rows:    height,
		Rule:    ruleNb,
		Start:   firstLine,
		Render:  renderType,
	}
}
