package automata

import (
	"math/rand"
)

// GetCenteredLine
func (c CellularAutomata) GetCenteredLine() []int {
	cells := make([]int, c.Columns)

	for i := range cells {
		cells[i] = 0
	}
	// Fill intermediates colors like this [0001234321000]
	for i := 0; i < c.States; i++ {
		offset := i - 1
		middle := c.Columns / 2
		cells[middle+offset] = i
		cells[middle-offset] = i
	}

	return cells
}

// GetRandomLine
func (c CellularAutomata) GetRandomLine() []int {
	cells := make([]int, c.Columns)

	for i := range cells {
		cells[i] = rand.Intn(c.States)
	}

	return cells
}

// GetCustomLine TODO get from command, for now get every odd index
func (c CellularAutomata) GetCustomLine() []int {
	cells := make([]int, c.Columns)

	for i := range cells {
		cells[i] = 2 % i
	}

	return cells
}

// GetWordLine TODO get word from command
// word:
// /!\ Only works with RenderAsText()
func (c CellularAutomata) GetWordLine( /* word string */ ) []int {

	word := [...]int{1, 2, 3, 4}

	cells := make([]int, c.Columns)

	for i := range cells {
		cells[i] = word[i%len(word)]
	}

	return cells
}
