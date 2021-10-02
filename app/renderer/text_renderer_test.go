package renderer

import (
	"strings"
	"testing"
)

// GenerateText  Prints the cellular automata to the console
func TestGenerateText(t *testing.T) {
	metadata := map[string]int{
		"states": 2,
		"order":  1,
		"rule":   123,
	}
	actual := GenerateText([][]int{
		{1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3},
		{0, 1, 2, 3, 4},
	}, metadata)
	expected :=
		`
		 ╔═════════════════════════╗
		 ║ States: 2  Order: 1     ║
		 ║ Rule:          123      ║
		 ╚═════════════════════════╝
    █
█████
▒▒▒▒▒
▓▓▓▓▓
█ ▒▓░
`
	actual = strings.ReplaceAll(actual, "\n", "")
	expected = strings.ReplaceAll(expected, "\n", "")
	// if !reflect.DeepEqual(actual, expected) {
	if actual != expected {
		t.Errorf("Text Renderer error: Got %v, expected %v.", actual, expected)
	}
}
