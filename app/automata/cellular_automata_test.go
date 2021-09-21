package automata

import (
	"reflect"
	"testing"
)

func TestCellularAuto2states1order(t *testing.T) {
	actual := CellularAutomata{
		States:     2,
		Order:      1,
		Columns:    5,
		Rows:       5,
		RuleNumber: 110,
		FirstLine:  FirstLine{Sequence: []int{0, 0, 1, 0, 0}},
	}.GetMatrix()

	expected := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Rule array is incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestCellularAutoRectangle(t *testing.T) {
	firstLine := FirstLine{GetCenteredLine(10, 2)}
	actual := CellularAutomata{2, 1, 10, 3, 110, firstLine}.GetMatrix()

	expected := [][]int{
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Rule array is incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestCellularAuto3states2ndOrder(t *testing.T) {
	firstLine := FirstLine{GetCenteredLine(8, 3)}
	actual := CellularAutomata{3, 2, 8, 8, 145194, firstLine}.GetMatrix()

	expected := [][]int{
		{0, 0, 0, 2, 1, 2, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 2, 1, 1, 0, 0, 0, 0},
		{1, 0, 0, 0, 1, 0, 0, 0},
		{1, 0, 0, 0, 1, 1, 0, 2},
		{0, 1, 0, 2, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0},
		{2, 0, 2, 0, 1, 0, 0, 0},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Rule array is incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestCellularAutoImageName(t *testing.T) {
	firstLine := FirstLine{GetCenteredLine(3, 2)}
	actual := CellularAutomata{3, 2, 3, 3, 123, firstLine}.GetMetadata()

	expected := map[string]int{
		"states": 3,
		"order":  2,
		"rule":   123,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected metadata to be %v, got: %v.", expected, actual)
	}
}
