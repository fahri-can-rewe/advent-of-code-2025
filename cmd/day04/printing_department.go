package main

import (
	"strings"
)

func parseInput(input string) [][]string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var result [][]string

	for _, line := range lines {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		result = append(result, row)
	}
	return result
}

func countForkliftAccess(grid [][]string) int {
	totalX := 0
	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == "@" {
				if countNeighbors(grid, row, col) < 4 {
					totalX++
				}
			}
		}
	}
	return totalX
}

func countNeighbors(grid [][]string, row, column int) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	for deltaRow := -1; deltaRow <= 1; deltaRow++ {
		for deltaCol := -1; deltaCol <= 1; deltaCol++ {
			if deltaRow == 0 && deltaCol == 0 {
				continue
			}
			neighborRow, neighborCol := row+deltaRow, column+deltaCol
			if neighborRow >= 0 && neighborRow < rows && neighborCol >= 0 && neighborCol < cols {
				if grid[neighborRow][neighborCol] == "@" {
					count++
				}
			}
		}
	}
	return count
}
