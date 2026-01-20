package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input1, err := util.ReadInput("cmd/day11/input1.txt")
	input2, err := util.ReadInput("cmd/day11/input2.txt")

	if err != nil {
		log.Fatal(err)
	}

	racksP1 := parseInput(input1)
	racksP2 := parseInput(input2)

	p1 := useDFS(racksP1, false)
	p2 := useDFS(racksP2, true)

	fmt.Println("Part 1 amount different paths to out:", p1)
	fmt.Println("Part 2 amount paths visit dac and fft:", p2)
}
