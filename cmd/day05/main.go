package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("cmd/day05/input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	ranges, ids := parseInput(input)
	p1 := countValidIDs(ids, ranges)
	p2 := countValidIDsInRange(ranges)

	fmt.Println("Part 1 number of fresh IDs ", p1)
	fmt.Println("Part 2 ingredient ID ranges ", p2)
}
