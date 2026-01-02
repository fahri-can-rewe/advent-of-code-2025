package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("../advent-of-code-2025/cmd/day02/input3.txt")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("The sum of invalid IDs for Part 1 is: ", SumInvalidIDs(input, true))
	fmt.Println("The sum of invalid IDs for Part 2 is: ", SumInvalidIDs(input, false))
}
