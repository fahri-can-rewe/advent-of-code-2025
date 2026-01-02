package main

import (
	"fmt"
	"log"

	"github.com/fahri-can-rewe/advent-of-code-2025/internal/util"
)

func main() {
	input, err := util.ReadInput("../advent-of-code-2025/cmd/day03/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day03: ", input)
}
