package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	Target  []bool
	Buttons [][]int
}

func parseInput(input string) []Machine {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var machines []Machine

	reTarget := regexp.MustCompile(`\[([.#]+)\]`)
	reButtons := regexp.MustCompile(`\(([\d,]+)\)`)

	for _, line := range lines {
		if line == "" {
			continue
		}
		targetMatch := reTarget.FindStringSubmatch(line)
		targetStr := targetMatch[1]
		target := make([]bool, len(targetStr))
		for i, char := range targetStr {
			if char == '#' {
				target[i] = true
			}
		}

		buttonMatches := reButtons.FindAllStringSubmatch(line, -1)
		buttons := make([][]int, len(buttonMatches))
		for i, match := range buttonMatches {
			nums := strings.Split(match[1], ",")
			for _, numStr := range nums {
				numStr = strings.TrimSpace(numStr)
				num, _ := strconv.Atoi(numStr)
				buttons[i] = append(buttons[i], num)
			}
		}

		machines = append(machines, Machine{
			Target:  target,
			Buttons: buttons,
		})
	}

	return machines
}

// findMinPresses
// Since we want the fewest presses, and each button press is binary (0 or 1),
// if numButtons is large, bitmasking 2^numButtons might be slow.
// Let's check if numButtons is reasonable.
// Max numButtons in input1.txt seems to be around 13 (line 110 has 13 buttons).
// 2^13 = 8192, which is very small.
// However, if it's larger, we'd need BFS or Gaussian elimination.
func findMinPresses(machine Machine) int {
	numLights := len(machine.Target)
	numButtons := len(machine.Buttons)

	minPresses := -1

	// Using bitmask
	limit := 1 << numButtons
	for i := 0; i < limit; i++ {
		state := make([]bool, numLights)
		presses := 0
		for b := 0; b < numButtons; b++ {
			if (i>>b)&1 == 1 {
				presses++
				for _, light := range machine.Buttons[b] {
					if light < numLights {
						state[light] = !state[light]
					}
				}
			}
		}
		minPresses = checkForMatch(machine, numLights, state, minPresses, presses)
	}

	return minPresses
}

func checkForMatch(machine Machine, numLights int, state []bool, minPresses int, presses int) int {
	match := true
	for l := 0; l < numLights; l++ {
		if state[l] != machine.Target[l] {
			match = false
			break
		}
	}
	if match {
		if minPresses == -1 || presses < minPresses {
			minPresses = presses
		}
	}
	return minPresses
}

func countBtnPress(input string) int {
	machines := parseInput(input)
	total := 0
	for _, m := range machines {
		res := findMinPresses(m)
		if res != -1 {
			total += res
		}
	}
	return total
}
