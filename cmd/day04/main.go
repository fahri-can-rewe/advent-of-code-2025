package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("cmd/day04/input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := parseInput(input)
	resPart1 := countForkliftAccess(grid)
	resPart2 := delAsMuchPaperAsPossible(grid)

	fmt.Println("Part 1 accessible rolls of paper: ", resPart1)
	fmt.Println("Part 2 rolls of paper which can be removed: ", resPart2)
}
