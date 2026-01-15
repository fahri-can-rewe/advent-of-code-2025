package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("cmd/day08/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	connections := 1000
	points := parseInput(input)
	p1 := connectJunctionBoxes(points, connections)
	fmt.Println("Part 1 size of three largest circuits:", p1)
	p2 := multiplyXCoordLastTwoJB(points)
	fmt.Println("Part 2 last connection X product:", p2)
}
