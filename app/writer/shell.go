package writer

import (
	"github.com/eliseduverdier/cellular-automata-go/app"
	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
)

func RenderShell() {
	params := parameters.GetFromShell()
	app.Render(params)
}
