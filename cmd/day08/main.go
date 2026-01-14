package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("cmd/day08/input_sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	points := parseInput(input)
	p1 := solve(points)
	fmt.Println("Part 1 three largest circuits ", p1)
}
