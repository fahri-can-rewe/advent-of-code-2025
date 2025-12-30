package main

import (
	"advent-of-code-2025/internal/util"
	"fmt"
	"log"
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

func main() {
	input, err := util.ReadInput("../advent-of-code-2025/cmd/day01/input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The password for Part1 is: ", moveDialP1(generateInstructions(input)))
	fmt.Println("The password for Part2 is: ", moveDialP2(generateInstructions(input)))
}

func generateInstructions(input string) []Rotator {
	instructions := make([]Rotator, 0, 100)
	regex := regexp.MustCompile(`([A-Z])(\d+)`)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		if matches != nil {
			key := matches[1]
			value, _ := strconv.Atoi(matches[2])
			instructions = append(instructions, Rotator{key, value})
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
