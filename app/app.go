package app

import (
	"github.com/eliseduverdier/cellular-automata-go/app/automata"
	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
	"github.com/eliseduverdier/cellular-automata-go/app/renderer"
)

func Render(params parameters.Parameters) {

	automata := automata.CellularAutomata{
		States:         params.States,
		Order:          params.Order,
		Columns:        params.Columns,
		Rows:           params.Rows,
		HasRandomStart: params.RandomStart,
		RuleNumber:     params.Rule,
	}

	switch params.Render {
	case "image":
		renderer.GenerateImage(automata.GetMatrix(), automata.GetMetadata())
	case "text":
		renderer.GenerateText(automata.GetMatrix())
	}
}
