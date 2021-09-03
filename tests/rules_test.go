package main

import (
	"reflect"
	"testing"

	"github.com/eliseduverdier/cellular-automata-go/app/automata"
)

func TestRule2states1order(t *testing.T) {
	actual := automata.RuleToArray(2, 1, 30)
	expected := []int{0, 1, 1, 1, 1, 0, 0, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Rule array is incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestRule3states2order(t *testing.T) {
	actual := automata.RuleToArray(3, 2, 12345)
	expected := []int{0, 2, 0, 1, 2, 2, 1, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Rule array is incorrect, got: %d, want: %d.", actual, expected)
	}
}
