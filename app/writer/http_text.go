package writer

import (
	"fmt"
	"net/http"

	"github.com/eliseduverdier/cellular-automata-go/app"
	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
)

func RenderTextPage(w http.ResponseWriter, req *http.Request) {
	image := app.RenderText(parameters.GetFromRequest(req))

	req.Header.Add("Application", "text/html")
	fmt.Fprintf(w,
		"<h1>Hello Cellular Automata</h1>"+
			"<pre>%s</pre>",
		image,
	)
}
