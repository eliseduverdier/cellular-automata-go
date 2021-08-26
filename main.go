package main

import (
	"flag"
	"github.com/eliseduverdier/cellular-automata-go/src/automata"
	"github.com/eliseduverdier/cellular-automata-go/src/renderer"

)

func main() {
	states := flag.Int("s", 2, "the number of states")
	order := flag.Int("o", 1, "the order (1 or 2)")
	columns := flag.Int("w", 100, "the number of columns")
	rows := flag.Int("h", 100, "the number of rows")
	rule := flag.Int("r", 73, "the rule number")
	randomStart := flag.Bool("random", true, "if the first line is random")

    automata := automata.CellularAutomata{
        States: *states,
        Order: *order,
        Columns: *columns,
        Rows: *rows,
        HasRandomStart: *randomStart,
        RuleNumber: *rule,
    }

    renderer.GenerateImage(automata.GetMatrix(), automata.GetImageName())
}