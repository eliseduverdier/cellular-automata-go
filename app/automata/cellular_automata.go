package automata

import (
	"math/rand"
	"strconv"
)

// CellularAutomata  The cellular automata structure
type CellularAutomata struct {
	States         int
	Order          int
	Columns        int
	Rows           int
	HasRandomStart bool
	RuleNumber     int
}

// GetMatrix  Builds the matrix line after line
func (c CellularAutomata) GetMatrix() [][]int {
	matrix := make([][]int, c.Rows)
	matrix[0] = c.getFirstLine()
	//for i:= range matrix {
	for i := 0; i < c.Rows-1; i++ {
		matrix[i+1] = c.getNextLine(matrix, i)
	}
	return matrix
}

// GetImageName  Returns the properties of the automata (states, order, and rule number)
func (c CellularAutomata) GetMetadata() map[string]int {
	return map[string]int{
		"states": c.States,
		"order":  c.Order,
		"rule":   c.RuleNumber,
	}
}

func (c CellularAutomata) getFirstLine() []int {
	cells := make([]int, c.Columns)

	if c.HasRandomStart {
		for i := range cells {
			cells[i] = rand.Intn(c.States)
		}
	} else {
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
	}

	return cells
}

func (c CellularAutomata) getNextLine(matrix [][]int, currentLineIndex int) []int {
	newLine := make([]int, c.Columns)

	// fmt.Println("---")
	for i := 0; i < c.Columns; i++ {
		newCellValue := c.newCell(i, currentLineIndex, matrix)
		newLine[i] = newCellValue
	}

	return newLine
}

func (c CellularAutomata) newCell(position int, currentLineIndex int, matrix [][]int) int {

	// fmt.Println(fmt.Sprintf("pos: %d, line: %d", position, currentLineIndex))
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

	ruleArray := RuleToArray(c.States, c.Order, c.RuleNumber)

	res := int(ruleArray[index])
	return res
}
