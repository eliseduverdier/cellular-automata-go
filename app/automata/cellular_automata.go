package automata

import (
	"strconv"
)

// CellularAutomata The cellular automata structure
type CellularAutomata struct {
	States     int
	Order      int
	Columns    int
	Rows       int
	RuleNumber int64
	FirstLine  FirstLine
}

// GetMatrix Builds the matrix line after line
func (c CellularAutomata) GetMatrix() [][]int {
	ruleArray := RuleToArray(c.States, c.Order, c.RuleNumber)

	matrix := make([][]int, c.Rows)
	matrix[0] = c.FirstLine.Sequence
	for i := 0; i < c.Rows-1; i++ {
		matrix[i+1] = c.getNextLine(matrix, i, ruleArray)
	}
	return matrix
}

// GetImageName Returns the properties of the automata (states, order, and rule number)
func (c CellularAutomata) GetMetadata() map[string]int {
	return map[string]int{
		"states": c.States,
		"order":  c.Order,
		"rule":   int(c.RuleNumber), // risk of being truncated...
	}
}

func (c CellularAutomata) getNextLine(matrix [][]int, currentLineIndex int, ruleArray []int) []int {
	newLine := make([]int, c.Columns)

	for i := 0; i < c.Columns; i++ {
		newCellValue := c.newCell(i, currentLineIndex, matrix, ruleArray)
		newLine[i] = newCellValue
	}

	return newLine
}

func (c CellularAutomata) newCell(position int, currentLineIndex int, matrix [][]int, ruleArray []int) int {

	sumOfBaseCells := matrix[currentLineIndex][position] * 10

	// handle diagonals cells on the sides: loop to the other side
	if position == 0 { // first
		sumOfBaseCells += matrix[currentLineIndex][c.Columns-1]*100 + matrix[currentLineIndex][1]
	} else if position == c.Columns-1 { // last
		sumOfBaseCells += matrix[currentLineIndex][position-1]*100 + matrix[currentLineIndex][0]
	} else {
		sumOfBaseCells += matrix[currentLineIndex][position-1]*100 + matrix[currentLineIndex][position+1]
	}

	if c.Order > 1 && currentLineIndex > 1 { // 2nd order: also add center cell from n-2 line
		sumOfBaseCells += matrix[currentLineIndex-1][position] * 1000
	}

	index, _ := strconv.ParseInt(strconv.Itoa(sumOfBaseCells), c.States, 10)

	res := int(ruleArray[index])
	return res
}
