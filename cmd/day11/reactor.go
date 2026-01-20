package main

import "strings"

const start = "you"
const end = "out"

func parseInput(input string) map[string][]string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	racks := make(map[string][]string)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 2. Split into "Key" and "Remainder"
		parts := strings.SplitN(line, ":", 2)

		key := strings.TrimSpace(parts[0])
		outputs := strings.Fields(parts[1])
		racks[key] = outputs
	}
	return racks
}

func countPaths(racks map[string][]string, current string, memo map[string]int) int {
	if current == end {
		return 1
	}

	if val, isOk := memo[current]; isOk {
		return val
	}

	totalPaths := 0
	outputs, isOk := racks[current]
	if !isOk {
		memo[current] = 0
		return 0
	}

	for _, neighbor := range outputs {
		totalPaths += countPaths(racks, neighbor, memo)
	}

	memo[current] = totalPaths
	return totalPaths
}

func useDFS(racks map[string][]string) int {
	memo := make(map[string]int)
	return countPaths(racks, start, memo)
}
