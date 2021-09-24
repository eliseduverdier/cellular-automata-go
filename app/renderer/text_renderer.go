package renderer

import (
	"fmt"
)

// GenerateText  Prints the cellular automata to the console
func GenerateText(matrix [][]int, metadata map[string]int) string {
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

	string := GenerateTitle(metadata)
	string += "\n"

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			string += colors[matrix[x][y]]
		}
		string += "\n"
	}

	return string
}

func GenerateTitle(metadata map[string]int) string {
	return fmt.Sprintf(
		`
		 ╔═════════════════════════╗
		 ║ States: %d  Order: %d     ║
		 ║ Rule: %12d      ║
		 ╚═════════════════════════╝`,
		metadata["states"],
		metadata["order"],
		metadata["rule"],
	)
}
