package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("cmd/day10/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	btnPresses := countBtnPress(input)
	fmt.Println("Part 1 fewest indicator lights on all of the machines:", btnPresses)
}
