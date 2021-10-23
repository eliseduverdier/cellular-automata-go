package automata

import (
	"math/rand"
)

type FirstLine struct {
	Sequence []int
}

// GetCenteredLine
func GetCenteredLine(length int, states int) []int {
	cells := initZeroedLine(length)
	// Fill intermediates colors like this [0001234321000]
	for i := 0; i < states; i++ {
		offset := i - 1
		middle := length / 2
		cells[middle+offset] = i
		cells[middle-offset] = i
	}

	return cells
}

// GetPixelsOnRight
func GetPixelsOnRight(length int, states int) []int {
	cells := initZeroedLine(length)
	for i := 1; i < states; i++ {
		cells[length-i] = i
	}
	return cells
}

// GetPixelsOnLeft
func GetPixelsOnLeft(length int, states int) []int {
	cells := initZeroedLine(length)
	for i := 1; i < states; i++ {
		cells[i-1] = i
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

// GetPixelsOnRight
func initZeroedLine(length int) []int {
	cells := make([]int, length)

	for i := range cells {
		cells[i] = 0
	}

	return cells
}
