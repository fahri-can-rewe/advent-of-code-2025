package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	fmt.Println("Advent of Code 2025 - Day 2")
	input, err := util.ReadInput("../advent-of-code-2025/cmd/day02/input3.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID ranges:\n", input)
	fmt.Println("Sum: ", SumInvalidIDs(input))
}
