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

func sumUpJoltages(input string, isPartTwo bool) int {
	banks := parseInput(input)
	var sum int
	for _, row := range banks {
		if isPartTwo {
			idxBiggestNUm := findIdxOfBiggestNum(row, len(row)-totalDigits)
			arr := findRemainingDigits(idxBiggestNUm, row)
			fmt.Println(transformArrToInt(arr))
		} else {
			sum += findTwoLargestDigits(row)
		}
	}

	return sum
}

//2. check from the biggest number index how many digits are currently left
//3. check how many digits you have to eliminate to get the remaining 11 biggest digits
//4. save your 12-digit number and return it

func findIdxOfBiggestNum(nums []int, lastPos int) int {
	idxBiggest := 0
	biggestNum := nums[0]
	for i := 1; i < lastPos; i++ {
		if biggestNum < nums[i] {
			idxBiggest = i
			biggestNum = nums[i]
		}
	}
	//fmt.Printf("Biggest number: %d index: %d\n", biggestNum, idxBiggest)
	return idxBiggest
}

func findRemainingDigits(idxBiggest int, nums []int) []int {
	fmt.Println(nums)
	twelveDigitNum := make([]int, 0, 12)
	twelveDigitNum = append(twelveDigitNum, nums[idxBiggest])
	amountRemainingDigits := len(nums) - idxBiggest - 1
	amountDigitsToCancel := len(nums) - 12 - idxBiggest
	i := idxBiggest + 1
	lastIdx := i
	for amountDigitsToCancel > 0 {
		if i == len(nums)-1 {
			break
		}
		if nums[i] > nums[i+1] {
			twelveDigitNum = append(twelveDigitNum, nums[i])
			amountRemainingDigits--
			lastIdx = i
		} else if nums[i] > nums[i-1] {
			twelveDigitNum = append(twelveDigitNum, nums[i])
			amountRemainingDigits--
			lastIdx = i
		}
		i++
	}
	twelveDigitNum = append(twelveDigitNum, nums[lastIdx+1:]...)
	return twelveDigitNum
}

func transformArrToInt(nums []int) int64 {
	var result int64
	for _, num := range nums {
		result = result*10 + int64(num)
	}
	return result
}
