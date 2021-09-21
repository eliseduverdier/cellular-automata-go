package writer_shell

import (
	"github.com/eliseduverdier/cellular-automata-go/app"
	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
)

// RenderShell creates or display an automata from shell
func RenderShell() {
	params := parameters.GetFromShell()
	app.Render(params)
}
