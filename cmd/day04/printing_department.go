package main

import (
	"strings"
)

func parseInput(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var result [][]rune

	for _, line := range lines {
		result = append(result, []rune(line))
	}
	return result
}

func countForkliftAccess(grid [][]rune) int {
	const maxAllowedNeighbors = 4
	totalX := 0
	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '@' {
				if countNeighbors(grid, row, col) < maxAllowedNeighbors {
					totalX++
				}
			}
		}
	}
	return totalX
}

func countNeighbors(grid [][]rune, row, column int) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	for deltaRow := -1; deltaRow <= 1; deltaRow++ {
		for deltaCol := -1; deltaCol <= 1; deltaCol++ {
			if deltaRow == 0 && deltaCol == 0 {
				continue
			}
			neighborRow, neighborCol := row+deltaRow, column+deltaCol
			areCoordinatesWithinBound := neighborRow >= 0 && neighborRow < rows && neighborCol >= 0 && neighborCol < cols
			if areCoordinatesWithinBound {
				if grid[neighborRow][neighborCol] == '@' {
					count++
				}
			}
		}
	}
	return count
}
