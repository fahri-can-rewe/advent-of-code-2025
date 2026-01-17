package main

import (
	"math"
	"sort"
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

func useOnlyRedAndGreenTiles(coordinates [][]int) int64 {
	xs := make(map[int]bool)
	ys := make(map[int]bool)
	for _, p := range coordinates {
		xs[p[0]] = true
		xs[p[0]+1] = true
		ys[p[1]] = true
		ys[p[1]+1] = true
	}

	sortedXs := sortAndUnique(xs)
	sortedYs := sortAndUnique(ys)

	xMap := make(map[int]int)
	for i, x := range sortedXs {
		xMap[x] = i
	}
	yMap := make(map[int]int)
	for i, y := range sortedYs {
		yMap[y] = i
	}

	gridWidth := len(sortedXs)
	gridHeight := len(sortedYs)
	grid := make([][]int, gridHeight)
	for i := range grid {
		grid[i] = make([]int, gridWidth)
	}

	// 1. Mark the Boundary
	for i := 0; i < len(coordinates); i++ {
		p1 := coordinates[i]
		p2 := coordinates[(i+1)%len(coordinates)]

		x1, y1 := p1[0], p1[1]
		x2, y2 := p2[0], p2[1]

		if x1 == x2 {
			startY, endY := minNum(y1, y2), maxNum(y1, y2)
			xi := xMap[x1]
			for y := startY; y <= endY; y++ {
				grid[yMap[y]][xi] = 1
			}
		} else {
			startX, endX := minNum(x1, x2), maxNum(x1, x2)
			yi := yMap[y1]
			for x := startX; x <= endX; x++ {
				grid[yi][xMap[x]] = 1
			}
		}
	}

	// 2. Fill the Interior
	for yi := 0; yi < gridHeight; yi++ {
		inside := false
		for xi := 0; xi < gridWidth; xi++ {
			if isVerticalBoundaryCompressed(grid, xi, yi) {
				inside = !inside
			}
			if inside {
				grid[yi][xi] = 1
			}
		}
	}

	pref := make([][]int64, gridHeight+1)
	for i := range pref {
		pref[i] = make([]int64, gridWidth+1)
	}

	for yi := 0; yi < gridHeight; yi++ {
		for xi := 0; xi < gridWidth; xi++ {
			val := int64(0)
			if grid[yi][xi] == 1 {
				val = 1
			}
			pref[yi+1][xi+1] = val + pref[yi][xi+1] + pref[yi+1][xi] - pref[yi][xi]
		}
	}

	getSum := func(xi1, yi1, xi2, yi2 int) int64 {
		r1, c1 := minNum(yi1, yi2), minNum(xi1, xi2)
		r2, c2 := maxNum(yi1, yi2), maxNum(xi1, xi2)
		return pref[r2+1][c2+1] - pref[r1][c2+1] - pref[r2+1][c1] + pref[r1][c1]
	}

	maxArea := int64(0)
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			p1, p2 := coordinates[i], coordinates[j]
			xi1, yi1 := xMap[p1[0]], yMap[p1[1]]
			xi2, yi2 := xMap[p2[0]], yMap[p2[1]]

			expectedCompressedArea := int64(absNum(xi1-xi2)+1) * int64(absNum(yi1-yi2)+1)

			if getSum(xi1, yi1, xi2, yi2) == expectedCompressedArea {
				area := int64(absNum(p1[0]-p2[0])+1) * int64(absNum(p1[1]-p2[1])+1)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func isVerticalBoundaryCompressed(grid [][]int, xi, yi int) bool {
	if grid[yi][xi] == 0 {
		return false
	}
	if yi > 0 && grid[yi-1][xi] == 1 {
		// This logic for orthogonal polygons: a vertical edge exists if we are on a vertical segment.
		// To be precise, we need to know if the segment (yi-1, yi) at xi is a boundary.
		// Since we marked all boundary points, this should work for scanline.
		return true
	}
	return false
}

func sortAndUnique(m map[int]bool) []int {
	res := make([]int, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absNum(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
