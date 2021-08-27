package renderer

import (
	"fmt"
)

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

	for x := 0; x < len(matrix[0]); x++ {
		for y := 0; y < len(matrix); y++ {
			fmt.Print(colors[matrix[x][y]])
		}
		fmt.Println("")
	}
}