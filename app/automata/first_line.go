package automata

import (
	"math/rand"
)

type FirstLine struct {
	Sequence []int
}

// GetCenteredLine
func GetCenteredLine(length int, states int) []int {
	cells := make([]int, length)

	for i := range cells {
		cells[i] = 0
	}
	// Fill intermediates colors like this [0001234321000]
	for i := 0; i < states; i++ {
		offset := i - 1
		middle := length / 2
		cells[middle+offset] = i
		cells[middle-offset] = i
	}

	return cells
}

// GetRandomLine
func GetRandomLine(length int, states int) []int {
	cells := make([]int, length)

	for i := range cells {
		cells[i] = rand.Intn(states)
	}

	return cells
}
