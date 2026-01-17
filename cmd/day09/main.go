package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("cmd/day09/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	coordinates := parseInput(input)
	area := findLargestRectangle(coordinates)
	fmt.Println("Part 1 largest area of any rectangle:", area)
	area2 := useOnlyRedAndGreenTiles(coordinates)
	fmt.Println("Part 2 largest area of any rectangle using only red and green tiles:", area2)
}
