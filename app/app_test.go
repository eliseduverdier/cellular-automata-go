package app

import (
	"testing"

	"github.com/eliseduverdier/cellular-automata-go/app/parameters"
)

func TestRenderImage(t *testing.T) {
	params := parameters.Parameters{
		States:  2,                                   // int
		Order:   1,                                   // int
		Columns: 10,                                  // int
		Rows:    10,                                  // int
		Start:   []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // []int
		Rule:    110,                                 // int
		Render:  "image",                             // string
	}

	actual := Render(params)
	expected := "s2-o1-r110"
	if actual != expected {
		t.Errorf("Test render error, expected %s, got %s.", expected, actual)
	}
}

func TestRenderText(t *testing.T) {
	params := parameters.Parameters{
		States:  2,                                   // int
		Order:   1,                                   // int
		Columns: 10,                                  // int
		Rows:    10,                                  // int
		Start:   []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // []int
		Rule:    110,                                 // int
		Render:  "text",                              // string
	}

	actual := Render(params)
	expected := ""
	if actual != expected {
		t.Errorf("Test render error, expected %s, got %s.", expected, actual)
	}
}
