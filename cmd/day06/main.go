package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("cmd/day06/input_sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	nums, operations := parseInput(input)
	p1 := sumUpCalculations(nums, operations)
	fmt.Println("Part 1 sum of all numbers: ", p1)
}
