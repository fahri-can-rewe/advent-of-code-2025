package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("../advent-of-code-2025/cmd/day01/input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The password for Part1 is: ", countDialPointsZero(generateInstructions(input)))
	fmt.Println("The password for Part2 is: ", countDialPointsAndPassedZero(generateInstructions(input)))
}
