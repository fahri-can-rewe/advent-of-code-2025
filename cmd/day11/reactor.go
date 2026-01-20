package main

import "strings"

const start = "you"
const end = "out"

type State struct {
	current string
	hasDac  bool
	hasFft  bool
}

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

func useDFS(racks map[string][]string, isPartTwo bool) int {
	memo := make(map[string]int)
	if isPartTwo {
		return countPathsWithTargets(racks, "svr", false, false, make(map[State]int))
	}
	return countPaths(racks, start, memo)
}

func countPathsWithTargets(racks map[string][]string, current string, hasDac, hasFft bool, memo map[State]int) int {
	// Update state if we just arrived at a target
	if current == "dac" {
		hasDac = true
	}
	if current == "fft" {
		hasFft = true
	}

	// Base Case
	if current == "out" {
		if hasDac && hasFft {
			return 1
		}
		return 0
	}

	state := State{current, hasDac, hasFft}
	if val, ok := memo[state]; ok {
		return val
	}

	total := 0
	for _, neighbor := range racks[current] {
		total += countPathsWithTargets(racks, neighbor, hasDac, hasFft, memo)
	}

	memo[state] = total
	return total
}
