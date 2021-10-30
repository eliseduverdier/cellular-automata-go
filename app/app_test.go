package app

import (
	"testing"

	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
)

func TestRenderImageName(t *testing.T) {
	params := parameters.Parameters{
		States:    2,
		Order:     1,
		Columns:   10,
		Rows:      10,
		PixelSize: 1,
		Start:     []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Rule:      110,
		Render:    "image",
	}

	actual := Render(params)
	expected := "s2-o1-r110"
	if actual != expected {
		t.Errorf("Test render error, expected %s, got %s.", expected, actual)
	}
}

func TestRenderTextImage(t *testing.T) {
	params := parameters.Parameters{
		States:    2,
		Order:     1,
		Columns:   10,
		Rows:      10,
		PixelSize: 1,
		Start:     []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Rule:      110,
		Render:    "text",
	}

	actual := Render(params)
	expected := ""
	if actual != expected {
		t.Errorf("Test render error, expected %s, got %s.", expected, actual)
	}
}

func TestRenderImage(t *testing.T) {
	params := parameters.Parameters{2, 1, 10, 10, 1, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 110, "image"}

	_, actualName := RenderImage(params)
	expectedName := "s2-o1-r110"
	if expectedName != actualName {
		t.Errorf("Test render error, expected name %s, got %s.", expectedName, actualName)
	}
}

func TestRenderText(t *testing.T) {
	params := parameters.Parameters{2, 1, 10, 9, 1, []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 0}, 110, "text"}

	actual := RenderText(params)
	expected :=
		`
		 ╔═════════════════════════╗
		 ║ States: 2  Order: 1     ║
		 ║ Rule:          110      ║
		 ╚═════════════════════════╝
████████ █
███████  █
██████   █
█████  █ █
████     █
███  ███ █
██   ██  █
█  █ █   █
       █ █
`
	if actual != expected {
		t.Errorf("Test render error, expected %s, got %s.", expected, actual)
	}
}
