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
		initialState := State{current: "svr", hasDac: false, hasFft: false}
		return countPathsWithTargets(racks, initialState, make(map[State]int))
	}
	return countPaths(racks, start, memo)
}

func countPathsWithTargets(racks map[string][]string, state State, memo map[State]int) int {
	// 1. Update the state if the current node is a target
	if state.current == "dac" {
		state.hasDac = true
	}
	if state.current == "fft" {
		state.hasFft = true
	}

	// 2. Base Case: Reached the end
	if state.current == "out" {
		if state.hasDac && state.hasFft {
			return 1
		}
		return 0
	}

	// 3. Memoization check
	if val, ok := memo[state]; ok {
		return val
	}

	total, isDone := exploreRecursive(racks, state, memo)
	if isDone {
		return 0
	}

	// 5. Store and return
	memo[state] = total
	return total
}

func exploreRecursive(racks map[string][]string, state State, memo map[State]int) (int, bool) {
	// 4. Recursive exploration
	total := 0
	outputs, exists := racks[state.current]
	if !exists {
		memo[state] = 0
		return 0, true
	}

	for _, neighbor := range outputs {
		// Create a new state for the neighbor while preserving the current hasDac/hasFft flags
		nextState := State{
			current: neighbor,
			hasDac:  state.hasDac,
			hasFft:  state.hasFft,
		}
		total += countPathsWithTargets(racks, nextState, memo)
	}
	return total, false
}
