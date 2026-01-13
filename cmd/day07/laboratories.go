package main

import "strings"

func parseInput(input string) [][]string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	diagram := make([][]string, len(lines))

	for i, line := range lines {
		//diagram[i] = []string(line)
		diagram[i] = strings.Split(line, "")
	}

	return diagram
}

func findS(diagram [][]string) int {
	var pos int
	for i := 0; i < len(diagram); i++ {
		for j := 0; j < len(diagram[i]); j++ {
			if diagram[i][j] == "S" {
				pos = j
			}
		}
	}
	return pos
}

func copyVal(diagram [][]string) [][]string {
	beams := make([][]string, len(diagram))

	for i := range diagram {
		// Create the inner slice with the same length
		beams[i] = make([]string, len(diagram[i]))
		// Copy the contents from diagram[i] to beams[i]
		copy(beams[i], diagram[i])
	}
	return beams
}

func countBeamSplitters(diagram [][]string) int {
	beams := copyVal(diagram)
	sPos := findS(diagram)
	counter := 0
	lastRow := len(diagram) - 1
	lastCol := len(diagram[0]) - 1

	for i := 2; i < len(diagram); i++ {
		for j := 0; j <= lastCol; j++ {
			if diagram[i][j] == "^" && i == 2 && j == sPos {
				beams[i+1][j-1] = "|"
				beams[i+1][j+1] = "|"
				counter++
			}
			if diagram[i][j] == "^" && beams[i-1][j] == "|" {
				if j != 0 && i < lastRow {
					beams[i+1][j-1] = "|"
				}
				if j != lastCol && i < lastRow {
					beams[i+1][j+1] = "|"
				}
				counter++
			}
		}
	}
	return counter
}
