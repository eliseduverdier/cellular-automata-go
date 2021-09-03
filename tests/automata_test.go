package main

import (
	"reflect"
	"testing"

	"github.com/eliseduverdier/cellular-automata-go/src/automata"
)

func TestCellularAuto2states1order(t *testing.T) {
	actual := automata.CellularAutomata{
		States:         2,
		Order:          1,
		Columns:        5,
		Rows:           5,
		HasRandomStart: false,
		RuleNumber:     110,
	}.GetMatrix()

	expected := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Rule array is incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestCellularAutoRectangle(t *testing.T) {
	actual := automata.CellularAutomata{2, 1, 10, 3, false, 110}.GetMatrix()

	expected := [][]int{
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Rule array is incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestCellularAuto3states2ndOrder(t *testing.T) {
	actual := automata.CellularAutomata{3, 2, 8, 8, false, 145194}.GetMatrix()

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
		t.Errorf("Rule array is incorrect, got: %d, want: %d.", actual, expected)
	}
}
func TestCellularAutoRandom(t *testing.T) {
	actual := automata.CellularAutomata{3, 2, 10, 1, true, 0}.GetMatrix()

	notExpected := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	if reflect.DeepEqual(actual, notExpected) {
		t.Errorf("Expected random to be random, got: %d.", actual)
	}
}

func TestCellularAutoImageName(t *testing.T) {
	actual := automata.CellularAutomata{3, 2, 3, 3, false, 123}.GetImageName()

	expected := "s3-o2-r123"

	if actual != expected {
		t.Errorf("Expected filename to be %s, got: %s.", expected, actual)
	}
}
