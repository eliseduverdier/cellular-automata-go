package automata

import (
	"reflect"
	"testing"
)

func TestRule2states1order(t *testing.T) {
	actual := RuleToArray(2, 1, 30)
	expected := []int{0, 1, 1, 1, 1, 0, 0, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Rule array is incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestRule3states2order(t *testing.T) {
	actual := RuleToArray(3, 2, 12345)
	expected := []int{0, 2, 0, 1, 2, 2, 1, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Rule array is incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestGetMaxStatesCombinaisons(t *testing.T) {
	// GetMaxStatesCombinaisons(2, 1) -> 8
	// GetMaxStatesCombinaisons(2, 2) -> 16
	// GetMaxStatesCombinaisons(3, 1) -> 27
	// GetMaxStatesCombinaisons(3, 2) -> 81
	nb := 2
	actual := GetMaxStatesCombinaisons(2, 1)
	expected := 8
	if actual != expected {
		t.Errorf("TestGetMaxStatesCombinaisons(%d) error, got: %d, expected: %d.", nb, actual, expected)
	}
}

func TestGetMaxRule(t *testing.T) {
	nb := 8
	// GetMaxRule(8) -> 256
	// GetMaxRule(16) -> 65536
	// GetMaxRule(27) -> 134217728
	// GetMaxRule(81) -> 2,417851639e24
	actual := GetMaxRule(nb)
	expected := float64(256)
	if actual != expected {
		t.Errorf("TestGetMaxRule(%v) error, got: %v, expected: %v.", nb, actual, expected)
	}
}
