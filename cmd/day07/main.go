package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("cmd/day07/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	diagram := parseInput(input)
	p1 := countBeamSplitters(diagram)
	fmt.Println("Part 1 tachyon beam split total of: ", p1)
}
