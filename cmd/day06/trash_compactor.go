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

func applyCephalopodMath(input string) int64 {
	lines := strings.Split(input, "\n")
	// Remove empty trailing line if exists
	if strings.TrimSpace(lines[len(lines)-1]) == "" && len(lines) > 0 {
		lines = lines[:len(lines)-1]
	}

	numRows := len(lines)
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// Ensure all lines have the same length by padding with spaces
	for i := range lines {
		if len(lines[i]) < maxLen {
			lines[i] = lines[i] + strings.Repeat(" ", maxLen-len(lines[i]))
		}
	}

	totalSum := int64(0)
	var currentProblemCols [][]string
	var currentProblemOps []string

	processProblem := func() {
		if len(currentProblemCols) == 0 {
			return
		}

		var nums []int64
		for _, colChars := range currentProblemCols {
			numStr := ""
			for _, char := range colChars {
				if char >= "0" && char <= "9" {
					numStr += char
				}
			}
			if numStr != "" {
				val, _ := strconv.ParseInt(numStr, 10, 64)
				nums = append(nums, val)
			}
		}

		op := ""
		for _, o := range currentProblemOps {
			if o == "+" || o == "*" {
				op = o
				break
			}
		}

		if len(nums) > 0 && op != "" {
			var res int64
			if op == "+" {
				res = 0
				for _, n := range nums {
					res += n
				}
			} else if op == "*" {
				res = 1
				for _, n := range nums {
					res *= n
				}
			}
			totalSum += res
		}

		currentProblemCols = nil
		currentProblemOps = nil
	}

	for j := maxLen - 1; j >= 0; j-- {
		isEmpty := true
		for i := 0; i < numRows; i++ {
			if lines[i][j] != ' ' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			processProblem()
		} else {
			var colChars []string
			// Digits are in all rows except the last one
			for i := 0; i < numRows-1; i++ {
				colChars = append(colChars, string(lines[i][j]))
			}
			currentProblemCols = append(currentProblemCols, colChars)

			// Operator is in the last row
			opChar := string(lines[numRows-1][j])
			if opChar != " " {
				currentProblemOps = append(currentProblemOps, opChar)
			}
		}
	}
	processProblem()

	return totalSum
}
