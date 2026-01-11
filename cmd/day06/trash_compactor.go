package main

import (
	"regexp"
	"strconv"
	"strings"
)

func extractNumbers(line string) []int {
	regex := regexp.MustCompile(`\d+`)
	matches := regex.FindAllString(line, -1)

	var nums []int
	for _, match := range matches {
		if val, err := strconv.Atoi(match); err == nil {
			nums = append(nums, val)
		}
	}
	return nums
}

func parseInput(input string) ([][]int, []string) {
	lines := strings.Split(input, "\n")
	var allNums [][]int
	var operations []string
	opRegex := regexp.MustCompile(`[+*]`)

	for _, line := range lines {
		// Extract numbers for this specific line
		nums := extractNumbers(line)
		if len(nums) > 0 {
			allNums = append(allNums, nums)
		}

		// Extract operations (this will find them even if mixed,
		// or specifically on the last line)
		ops := opRegex.FindAllString(line, -1)
		if len(ops) > 0 {
			operations = append(operations, ops...)
		}
	}
	return allNums, operations
}

func multiply(nums [][]int, col int) int64 {
	product := int64(1)
	for i := 0; i < len(nums); i++ {
		product *= int64(nums[i][col])
	}
	return product
}

func add(nums [][]int, col int) int64 {
	sum := int64(0)
	for i := 0; i < len(nums); i++ {
		sum += int64(nums[i][col])
	}
	return sum
}

func runOperation(nums [][]int, operations []string) int64 {
	sum := int64(0)
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "+":
			sum += add(nums, i)
		case "*":
			sum += multiply(nums, i)
		}
	}
	return sum
}
