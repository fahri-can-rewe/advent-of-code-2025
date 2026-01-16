package main

import (
	"math"
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	coordinates := make([][]int, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		xy := strings.Split(line, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		coordinates = append(coordinates, []int{x, y})
	}
	return coordinates
}

func findLargestRectangle(coordinates [][]int) int64 {
	maxArea := int64(0)
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			x := math.Abs(float64(coordinates[j][0]-coordinates[i][0])) + 1
			y := math.Abs(float64(coordinates[j][1]-coordinates[i][1])) + 1
			area := x * y
			recArea := int64(area)
			if recArea > maxArea {
				maxArea = recArea
			}
		}
	}
	return maxArea
}
