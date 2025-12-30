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

func moveDialP1(rotations []Rotator) int {
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

func moveDialP2(rotations []Rotator) int {
	dialPos := 50
	counter := 0

	for _, rotator := range rotations {
		if rotator.steps <= 0 {
			continue
		}

		var distToZero int
		if rotator.direction == "R" {
			// If at 0, the next 0 is 100 clicks away. Otherwise, it's 100 - pos.
			if dialPos == 0 {
				distToZero = allDialPositions
			} else {
				distToZero = allDialPositions - dialPos
			}

			// New position after rotation
			dialPos = (dialPos + rotator.steps) % allDialPositions
		} else {
			// Moving Left, the distance to zero is simply the current position.
			// If at 0, the next 0 is 100 clicks away.
			if dialPos == 0 {
				distToZero = allDialPositions
			} else {
				distToZero = dialPos
			}

			// New position after rotation (Go-safe modulo)
			dialPos = (dialPos - (rotator.steps % allDialPositions) + allDialPositions) % allDialPositions
		}

		// Logic: If the steps taken meet or exceed the distance to the first zero,
		// we count 1, then add 1 for every additional 100 steps.
		if rotator.steps >= distToZero {
			counter += 1 + (rotator.steps-distToZero)/allDialPositions
		}
	}

	return counter
}
