package main

import (
	"github.com/eliseduverdier/cellular-automata-go/src/automata"
	"github.com/eliseduverdier/cellular-automata-go/src/renderer"

)

func main() {
    matrix := automata.CellularAutomata{
        States: 2,
        Order: 1,
        Columns: 100,
        Rows: 100,
        HasRandomStart: false,
        RuleNumber: 110,
    }.GetMatrix()

    renderer.GenerateImage(matrix)
}