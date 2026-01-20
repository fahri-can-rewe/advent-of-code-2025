package main

import "strings"

const start = "you"
const end = "out"

type ServerRack struct {
	key     string
	outputs []string
}

type Device struct {
	id         int
	subDevices Queue[string]
}

func parseInput(input string) []ServerRack {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rack := make([]ServerRack, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		} // Skip empty lines

		// 2. Split into "Key" and "Remainder"
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue
		} // Guard against lines without a colon

		rack = append(rack, ServerRack{
			key:     strings.TrimSpace(parts[0]),
			outputs: strings.Fields(parts[1]),
		})
	}
	return rack
}

func canReach(racks []ServerRack, current string, visited map[string]bool) bool {
	// If we found the target
	if current == end {
		return true
	}

	// Mark current as visited to avoid cycles
	visited[current] = true

	// Find the current rack
	var currentRack *ServerRack
	for i := range racks {
		if racks[i].key == current {
			currentRack = &racks[i]
			break
		}
	}

	if currentRack == nil {
		return false
	}

	// Explore neighbors
	for _, neighbor := range currentRack.outputs {
		// Skip already visited nodes
		if visited[neighbor] {
			continue
		}

		if canReach(racks, neighbor, visited) {
			return true
		}
	}

	return false
}

func dfs(racks []ServerRack) int {
	steps := 0

	// Find the starting rack
	var startRack *ServerRack
	for i := range racks {
		if racks[i].key == start {
			startRack = &racks[i]
			break
		}
	}

	if startRack == nil {
		return 0
	}

	// For each immediate child, check if it can reach target
	for _, child := range startRack.outputs {
		visited := make(map[string]bool)
		if canReach(racks, child, visited) {
			steps++
		}
	}

	return steps
}
