package CellularAutomata

import "math"
import "strconv"

// Generates an associative array,
// where N cells with S states generate a new cell
// Basic cellular automata (2 states, 1st order), based on the 3 cells above.
// For example, for the rule 30 (11110 in base 2), the array map these values :
// current pattern             111  110  101  100  011  010  001  000
// new state for center cell    0    0    0    1    1    1    1    0

// Idem for (4 states/2nd order), where the array is larger
// 0001  0002  0003  0010  ...  3333
//  2     3     0     1          1
func ruleToArray(states int, order int, ruleNumber int) []int {
	// o cells with n possible states: n^o + 1
	lengthOfRuleNumber := math.Pow(float64(states), float64(order + 2)) + 1

	// toBaseN := sprintf('%0' . lengthOfRuleNumber . 's', base_convert(ruleNumber, 10, this->states))
	toBaseN := strconv.FormatInt(int64(ruleNumber), states)

	return array_reverse(str_split(strval(toBaseN)))
}