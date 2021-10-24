package automata

import (
	"reflect"
	"testing"
)

type Test struct {
	Actual   []int
	Expected []int
}

// GetLineFromStrategyTest
func TestGetCenteredLine(t *testing.T) {
	tests := []Test{
		{Actual: GetCenteredLine(10, 2), Expected: []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0}},
		{Actual: GetCenteredLine(9, 2), Expected: []int{0, 0, 0, 0, 1, 0, 0, 0, 0}},
		{Actual: GetCenteredLine(10, 5), Expected: []int{0, 0, 4, 3, 2, 1, 2, 3, 4, 0}},
		{Actual: GetPixelsOnLeft(10, 2), Expected: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{Actual: GetPixelsOnLeft(10, 6), Expected: []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}},
		{Actual: GetPixelsOnRight(10, 2), Expected: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{Actual: GetPixelsOnRight(10, 6), Expected: []int{0, 0, 0, 0, 0, 5, 4, 3, 2, 1}},
	}

	for i := range tests {
		if !reflect.DeepEqual(tests[i].Actual, tests[i].Expected) {
			t.Errorf("Centered first line is incorrect, got: %d, expected: %d.", tests[i].Actual, tests[i].Expected)
		}
	}
}

// GetRandomLineTest
func TestGetRandomLine(t *testing.T) {
	expectedLen := 10
	randomLine := GetRandomLine(expectedLen, 2)
	if len(randomLine) != expectedLen {
		t.Errorf("Random line is incorrect, got: %d, expected length of %d.", randomLine, expectedLen)
	}
}
