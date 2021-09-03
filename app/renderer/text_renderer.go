package renderer

import (
	"fmt"
)

// GenerateText  Prints the cellular automata to the console
func GenerateText(matrix [][]int) {

	colors := []string{
		"█",
		" ",
		"▒",
		"▓",
		"░",
		"▲",
		"❖",
		"✨",
		"❂",
		"❄",
	}

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			fmt.Print(colors[matrix[x][y]])
		}
		fmt.Println("")
	}
}
