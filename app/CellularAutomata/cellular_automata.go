package CellularAutomata

import "math/rand"
import "strconv"

type CellularAutomata struct {
	states int
	order int
    columns  int
    rows    int
    hasRandomStart bool
	ruleNumber int
	matrix [][]int
}

func (c CellularAutomata) getFirstLine() []int {
	cells := make([]int, c.columns)

    if c.hasRandomStart {
		for i := range cells {
            cells[i] = rand.Intn(c.states - 1)
        }
    } else {
		for i := range cells {
            cells[i] = 0
        }
        // Fill intermediates colors like this [0001234321000]
        for i := 0 ; i < c.states ; i++{
            offset := i - 1;
            middle := c.columns / 2
            cells[middle + offset] = i;
            cells[middle - offset] = i;
        }
    }

    return cells
}

func (c CellularAutomata) getNextLine(matrix [][]int, currentLineIndex int) []int {
	newLine := make([]int, c.columns)

	for i := 0; i < c.columns; i++ {
		newcellvalue := c.newCell(i, currentLineIndex, matrix);
		newLine = append(newLine, newcellvalue);
	}

	return newLine
}

func (c CellularAutomata) newCell(position int, currentLineIndex int, matrix [][]int) int {

	sumOfBaseCells := matrix[currentLineIndex][position] * 10

    // handle diagonals cells on the sides: loop to the other side
    if position == 0 { // first
        sumOfBaseCells += matrix[currentLineIndex][c.columns] * 100 + matrix[currentLineIndex][1]
    } else if position == c.columns { // last
        sumOfBaseCells += matrix[currentLineIndex][position - 1] * 100 + matrix[currentLineIndex][0]
    } else {
        sumOfBaseCells += matrix[currentLineIndex][position - 1] * 100 + matrix[currentLineIndex][position + 1]
    }

    if c.order == 2 && currentLineIndex > 1 { // 2nd order: also add center cell from n-2 line
        sumOfBaseCells += matrix[currentLineIndex - 1][position] * 1000
    }

    index := strconv.FormatInt(sumOfBaseCells, c.states)

	ruleArray := ruleToArray(c.ruleNumber);

    return int(ruleArray[index]);
}