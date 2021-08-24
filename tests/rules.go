/* TODO
Assert.equals (
	ruleToArray(2, 1, 30), // 2 states, 1st order, rule 30
    [0 1 1 1 1 0 0 0]
)

Assert.equals (
	ruleToArray(3, 2, 12345), // 2 states, 1st order, rule 30
	[0 2 0 1 2 2 1 2 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
)
*/
package main

import "testing"

// TODO tests/rules.go:15:8: cannot find package "../app/CellularAutomata/rule_generator.go" in:
import "../app/CellularAutomata/rule_generator.go"

func TestRule(t *testing.T) {
    actual := ruleToArray(2, 1, 30)
	expected = [8]int{0, 1, 1, 1, 1, 0, 0, 0}
    if actual != expected {
       t.Errorf("Rule array was incorrect, got: %d, want: %d.", actual, expected)
    }
}