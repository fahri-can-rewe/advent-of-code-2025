package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("cmd/day11/input_sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	racks := parseInput(input)
	steps := dfs(racks)
	fmt.Println("Part 1 amount different paths to out:", steps)
}
