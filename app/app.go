package app

import (
	"image"
	"path"
	"runtime"

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
		renderer.GenerateImage(automata.GetMatrix(), automata.GetMetadata(), GetImagePath())
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

func RenderImage(params parameters.Parameters) (*image.RGBA, string) {

	automata := automata.CellularAutomata{
		States:     params.States,
		Order:      params.Order,
		Columns:    params.Columns,
		Rows:       params.Rows,
		RuleNumber: params.Rule,
		FirstLine:  automata.FirstLine{params.Start},
	}

	return renderer.GenerateImage(automata.GetMatrix(), automata.GetMetadata(), GetImagePath())
}

func GetImagePath() string {
	_, currentFile, _, _ := runtime.Caller(0)
	imageDir := path.Join(currentFile, "/../../images")
	return imageDir
}
