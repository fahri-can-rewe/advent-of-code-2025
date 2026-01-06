package main

import (
	"fmt"
	"strconv"
	"strings"
)

const biggestDigit = 9

func parseInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var result [][]int

	for _, line := range lines {
		var row []int
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		result = append(result, row)
	}
	return result
}

func isHighestNum(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	lastIdx := len(nums) - 1
	if nums[0] == biggestDigit && nums[1] == biggestDigit {
		return true
	}
	if len(nums) > 2 {
		if nums[0] == biggestDigit && nums[lastIdx] == biggestDigit {
			return true
		}
		if nums[1] == biggestDigit && nums[lastIdx] == biggestDigit {
			return true
		}
	}
	return false
}

func findSecondDigit(idxBiggest int, nums []int) int {
	idxLastNum := 1
	startIdx := idxBiggest + 1
	for i := startIdx; i < len(nums); i++ {
		if idxLastNum < nums[i] {
			idxLastNum = nums[i]
			idxBiggest = i
		}
	}
	return idxLastNum
}

func findTwoLargestDigits(nums []int) int {
	const biggestNum = 99
	if len(nums) == 0 {
		return 0
	}
	if isHighestNum(nums) {
		return biggestNum
	}
	if len(nums) == 1 {
		return nums[0]
	}
	firstDigit := nums[0]
	secondDigit := 1
	lastIdx := len(nums) - 1
	idxBiggest := 0

	if firstDigit == biggestDigit {
		secondDigit = findSecondDigit(idxBiggest, nums)
	} else {
		for i := 1; i < lastIdx; i++ {
			if firstDigit < nums[i] {
				firstDigit = nums[i]
				idxBiggest = i
			}
		}
		secondDigit = findSecondDigit(idxBiggest, nums)
	}
	combinedNums := fmt.Sprintf("%d%d", firstDigit, secondDigit)
	result, _ := strconv.Atoi(combinedNums)
	return result
}

func sumUpJoltages(input string) int {
	banks := parseInput(input)
	var sum int
	for _, row := range banks {
		sum += findTwoLargestDigits(row)
	}
	return sum
}
