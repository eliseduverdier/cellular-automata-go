package automata

import "math"
import "strconv"

// Generates an associative array,
// where N cells with S states generate a new cell
// Basic cellular automata (2 states, 1st order), based on the 3 cells above.
// For example, for the rule 30 (11110 in base 2), the array map these values :
// current pattern          : 111 (7)  110 (6)  101 (5)  100 (4)  011 (3)  010 (2)  001 (1)  000 (0)
// new state for center cell:   0        0        0        1        1        1        1        0
func RuleToArray(states int, order int, ruleNumber int) []int {
	// o cells with n possible states: n^o + 1
	maxStates := int(math.Pow(float64(states), float64(order + 2)))

	ruleInBaseN := strconv.FormatInt(int64(ruleNumber), states)
	reversedRule := Reverse(ruleInBaseN)

	ruleArray := make([]int, maxStates)
	for i := 0 ; i < maxStates ; i++ {
		if i < len(reversedRule) {
			ruleArray[i] = reversedRule[i]
		} else {
			ruleArray[i] = 0
		}
	}

	return ruleArray
}

// s:      12110
// return: [0, 1, 1, 2, 1]
func Reverse(number string) []int {
	//s := string(number)
	result := make([]int, len(number))
	for pos, char := range number {
    i, _ := strconv.Atoi(string(char))
    result[len(number) - 1 - pos] = i
	}

	return result
}