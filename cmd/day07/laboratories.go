package main

import "strings"

func parseInput(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	diagram := make([][]rune, len(lines))

	for i, line := range lines {
		diagram[i] = []rune(line)
	}

	return diagram
}

func findS(diagram [][]rune) int {
	var pos int
	for i := 0; i < len(diagram); i++ {
		for j := 0; j < len(diagram[i]); j++ {
			if diagram[i][j] == 'S' {
				pos = j
			}
		}
	}
	return pos
}

func countBeamSplitters(diagram [][]rune) int {
	rows := len(diagram)
	cols := len(diagram[0])

	beamAt := make([][]bool, rows)
	for i := range beamAt {
		beamAt[i] = make([]bool, cols)
	}

	sCol := findS(diagram)
	// The beam starts at S and always moves downward.
	// We can process row by row.
	beamAt[0][sCol] = true

	splittersCounted := make(map[[2]int]bool)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !beamAt[r][c] {
				continue
			}

			checkForSplitters(diagram, r, c, splittersCounted, rows, beamAt, cols)
		}
	}

	return len(splittersCounted)
}

func checkForSplitters(diagram [][]rune, r int, c int, splittersCounted map[[2]int]bool, rows int, beamAt [][]bool, cols int) {
	// If it's a splitter, it stops the beam and emits two new ones
	if diagram[r][c] == '^' {
		splittersCounted[[2]int{r, c}] = true
		// New beams continue from immediate left and immediate right
		if r+1 < rows {
			if c > 0 {
				beamAt[r+1][c-1] = true
			}
			if c+1 < cols {
				beamAt[r+1][c+1] = true
			}
		}
	} else {
		// For empty space (.) or start (S) or already formed beam (|),
		// the beam continues downward.
		if r+1 < rows {
			beamAt[r+1][c] = true
		}
	}
}
