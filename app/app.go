package app

import (
	"github.com/eliseduverdier/cellular-automata-go/app/automata"
	"github.com/eliseduverdier/cellular-automata-go/app/renderer"
)

func Render(states int, order int, columns int, rows int, randomStart bool, rule int, renderAs string) {

	automata := automata.CellularAutomata{
		States:         states,
		Order:          order,
		Columns:        columns,
		Rows:           rows,
		HasRandomStart: randomStart,
		RuleNumber:     rule,
	}

	switch renderAs {
	case "image":
		renderer.GenerateImage(automata.GetMatrix(), automata.GetMetadata())
	case "text":
		renderer.GenerateText(automata.GetMatrix())
	}
}
