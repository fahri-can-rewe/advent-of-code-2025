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

func extractOperations(input string, operations []string) []string {
	opRegex := regexp.MustCompile(`[+*]`)
	ops := opRegex.FindAllString(input, -1)
	operations = append(operations, ops...)
	return operations
}

func parseInput(input string) ([][]int, []string) {
	lines := strings.Split(input, "\n")
	var allNums [][]int
	var operations []string
	operations = extractOperations(input, operations)

	for _, line := range lines {
		nums := extractNumbers(line)
		if len(nums) > 0 {
			allNums = append(allNums, nums)
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

func sumUpCalculations(nums [][]int, operations []string) int64 {
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
