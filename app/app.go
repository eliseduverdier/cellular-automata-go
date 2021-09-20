package app

import (
	"image"

	"github.com/eliseduverdier/cellular-automata-go/app/automata"
	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
	"github.com/eliseduverdier/cellular-automata-go/app/renderer"
)

func Render(params parameters.Parameters) string {

	automata := automata.CellularAutomata{
		States:     params.States,
		Order:      params.Order,
		Columns:    params.Columns,
		Rows:       params.Rows,
		RuleNumber: params.Rule,
		FirstLine:  automata.FirstLine{params.Start},
	}

	switch params.Render {
	case "image":
		renderer.GenerateImage(automata.GetMatrix(), automata.GetMetadata())
		return renderer.GetImageName(automata.GetMetadata())
	case "text":
		renderer.GenerateText(automata.GetMatrix(), automata.GetMetadata())
		return ""
	default:
		return "" // TODO, panic() and return type with err
	}
}

func RenderText(params parameters.Parameters) string {

	automata := automata.CellularAutomata{
		States:     params.States,
		Order:      params.Order,
		Columns:    params.Columns,
		Rows:       params.Rows,
		RuleNumber: params.Rule,
		FirstLine:  automata.FirstLine{params.Start},
	}

	return renderer.GenerateText(automata.GetMatrix(), automata.GetMetadata())
}

func RenderImage(params parameters.Parameters) *image.RGBA {

	automata := automata.CellularAutomata{
		States:     params.States,
		Order:      params.Order,
		Columns:    params.Columns,
		Rows:       params.Rows,
		RuleNumber: params.Rule,
		FirstLine:  automata.FirstLine{params.Start},
	}

	return renderer.GenerateImage(automata.GetMatrix(), automata.GetMetadata())
}
