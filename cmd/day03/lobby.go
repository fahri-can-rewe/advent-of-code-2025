package main

import (
	"fmt"
	"strconv"
	"strings"
)

const biggestDigit = 9
const totalDigits = 12

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
	secondDigit := 1
	idxBiggest := 0

	if nums[idxBiggest] == biggestDigit {
		secondDigit = findSecondDigit(idxBiggest, nums)
	} else {
		lastIdx := len(nums) - 1
		idxBiggest = findIdxOfBiggestNum(nums, lastIdx)
		secondDigit = findSecondDigit(idxBiggest, nums)
	}
	combinedNums := fmt.Sprintf("%d%d", nums[idxBiggest], secondDigit)
	result, _ := strconv.Atoi(combinedNums)
	return result
}

func sumUpJoltages(input string, isPartTwo bool) int64 {
	banks := parseInput(input)
	var sum int64
	for _, row := range banks {
		if len(row) == 0 {
			continue
		}
		if isPartTwo {
			idxBiggestNum := findIdxOfBiggestNum(row, len(row)-totalDigits+1)
			arr := findRemainingDigits(idxBiggestNum, row)
			sum += transformArrToInt(arr)
		} else {
			sum += int64(findTwoLargestDigits(row))
		}
	}

	return sum
}

func findIdxOfBiggestNum(nums []int, lastPos int) int {
	idxBiggest := 0
	biggestNum := nums[0]
	for i := 1; i < lastPos; i++ {
		if biggestNum < nums[i] {
			idxBiggest = i
			biggestNum = nums[i]
		}
	}
	return idxBiggest
}

func findRemainingDigits(idxBiggest int, nums []int) []int {
	remaining := nums[idxBiggest:]
	toRemove := len(remaining) - totalDigits

	if toRemove <= 0 {
		return remaining
	}

	var stack []int
	removed := 0
	for _, num := range remaining {
		for removed < toRemove && len(stack) > 1 && stack[len(stack)-1] < num {
			stack = stack[:len(stack)-1]
			removed++
		}
		stack = append(stack, num)
	}

	// If we still need to remove digits, remove from the end
	for removed < toRemove {
		stack = stack[:len(stack)-1]
		removed++
	}
	return stack
}

func transformArrToInt(nums []int) int64 {
	var result int64
	for _, num := range nums {
		result = result*10 + int64(num)
	}
	return result
}
