package main


import (
    "reflect"
    "testing"
    "github.com/eliseduverdier/cellular-automata-go/src/automata"
)

func TestCellularAuto2states1order(t *testing.T) {
    actual := automata.CellularAutomata{
        States: 2,
        Order: 1,
        Columns: 5,
        Rows: 5,
        HasRandomStart: false,
        RuleNumber: 110,
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
