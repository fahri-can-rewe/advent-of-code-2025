package main

import (
	"regexp"
	"strconv"
	"strings"
)

const lowestPoint = 0
const highestPoint = 99
const allDialPositions = 100

type Rotator struct {
	direction string
	steps     int
}

func generateInstructions(input string) []Rotator {
	instructions := make([]Rotator, 0, 100)
	regex := regexp.MustCompile(`([A-Z])(\d+)`)
	lines := strings.Split(input, "\n")
	const charIdx = 1
	const stepsIdx = 2
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		if matches != nil {
			letter := matches[charIdx]
			amount, _ := strconv.Atoi(matches[stepsIdx])
			instructions = append(instructions, Rotator{letter, amount})
		}
	}
	return instructions
}

func countDialPointsZero(rotations []Rotator) int {
	dialPos := 50
	counter := 0
	for _, rotator := range rotations {
		twoDigitNum := rotator.steps % allDialPositions
		if rotator.direction == "R" {
			if (twoDigitNum + dialPos) > highestPoint {
				dialPos = (dialPos + twoDigitNum) - 1
				dialPos = dialPos - highestPoint
			} else {
				dialPos += twoDigitNum
			}
		} else {
			if (dialPos - twoDigitNum) < lowestPoint {
				dialPos = (dialPos - twoDigitNum) + 1
				dialPos = highestPoint + dialPos
			} else {
				dialPos -= twoDigitNum
			}
		}
		if dialPos == lowestPoint {
			counter++
		}
	}
	return counter
}

func countDialPointsAndPassedZero(rotations []Rotator) int {
	dialPos := 50
	counter := 0
	for _, rotator := range rotations {
		if rotator.direction == "R" {
			counter += (dialPos + rotator.steps) / allDialPositions
			dialPos = (dialPos + rotator.steps) % allDialPositions
		} else {
			// If we start at 0, we only hit 0 again after a full 100 steps
			startPos := dialPos
			if dialPos == 0 {
				startPos = allDialPositions
			}
			// Check if we move far enough to reach or pass the '0' mark.
			// Note: if we start at 0, 'startPos' is 100, meaning we only hit
			// zero again if we complete a full rotation.
			if rotator.steps >= startPos {
				// We count 1 for the first time we hit/pass zero,
				// then add 1 for every full 100-step rotation thereafter.
				counter += 1 + (rotator.steps-startPos)/allDialPositions
			}

			// Update the dial position.
			// We use (steps % 100) to find the relative movement, subtract it from
			// the current position, and add 100 before taking the modulo again
			// to ensure the result is always a positive index between 0 and 99.
			dialPos = (dialPos - (rotator.steps % allDialPositions) + allDialPositions) % allDialPositions
		}
	}
	return counter
}
