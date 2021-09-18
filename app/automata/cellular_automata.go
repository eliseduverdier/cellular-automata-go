package automata

import (
	"strconv"
)

// CellularAutomata  The cellular automata structure
type CellularAutomata struct {
	States        int
	Order         int
	Columns       int
	Rows          int
	FirstLineType string
	RuleNumber    int
}

// var ruleArray []int

// GetMatrix  Builds the matrix line after line
func (c CellularAutomata) GetMatrix( /* options []string */ ) [][]int {
	ruleArray := RuleToArray(c.States, c.Order, c.RuleNumber)

	matrix := make([][]int, c.Rows)
	matrix[0] = c.getFirstLine(c.FirstLineType /* , options */)
	//for i:= range matrix {
	for i := 0; i < c.Rows-1; i++ {
		matrix[i+1] = c.getNextLine(matrix, i, ruleArray)
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

// getFirstLine TODO strategy type ?
func (c CellularAutomata) getFirstLine(strategy string /* , options []string */) []int {
	switch strategy {
	case "centered":
		return c.GetCenteredLine()
	case "custom":
		return c.GetCustomLine() // <- TODO get array from custom line
	case "word":
		return c.GetWordLine( /* options */ ) // <- TODO get word
	default: //case "random":
		return c.GetRandomLine()
	}
}

func (c CellularAutomata) getNextLine(matrix [][]int, currentLineIndex int, ruleArray []int) []int {
	newLine := make([]int, c.Columns)

	// fmt.Println("---")
	for i := 0; i < c.Columns; i++ {
		newCellValue := c.newCell(i, currentLineIndex, matrix, ruleArray)
		newLine[i] = newCellValue
	}

	return newLine
}

func (c CellularAutomata) newCell(position int, currentLineIndex int, matrix [][]int, ruleArray []int) int {

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

	// TODO: move that! no need to call the function each time !
	// ruleArray := RuleToArray(c.States, c.Order, c.RuleNumber)

	res := int(ruleArray[index])
	return res
}
