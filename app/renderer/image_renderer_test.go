package renderer

import (
	_ "fmt"
	_ "os"
	"testing"

	_ "github.com/eliseduverdier/cellular-automata-go/app/automata"
)

func TestImage(t *testing.T) {
	/* automata := automata.CellularAutomata{
	        States: 2,
	        Order: 1,
	        Columns: 5,
	        Rows: 5,
	        RuleNumber: 123,
	        FirstLine: automata.FirstLine({[1,1,1,0,0]}),
	    }
		renderer.GenerateImage(automata.GetMatrix(), automata.GetMetadata())

		currentTestDir, _ := t.TempDir() // added in go1.15 :()
		expectedPath := currentTestDir + "/../images/s2-o1-r123.png"
		_, err := os.Open(expectedPath)

		if err != nil {
			t.Errorf("Image '%s' was not found", expectedPath)
		} */
}
