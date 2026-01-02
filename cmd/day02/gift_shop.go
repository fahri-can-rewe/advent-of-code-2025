package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type IDRange struct {
	start int64
	end   int64
}

func parseIDRanges(input string) ([]IDRange, error) {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")
	ranges := make([]IDRange, 0, len(parts))
	regex := regexp.MustCompile(`^(\d+)-(\d+)$`)

	for _, part := range parts {
		match := regex.FindStringSubmatch(strings.TrimSpace(part))
		if match == nil {
			return nil, errors.New("invalid range format: " + part)
		}
		start, _ := strconv.ParseInt(match[1], 10, 64)
		end, _ := strconv.ParseInt(match[2], 10, 64)
		if start > end {
			return nil, errors.New("invalid range: start cannot be greater than end")
		}
		ranges = append(ranges, IDRange{start: start, end: end})
	}
	return ranges, nil
}

func areNumbersRepeating(num int64) bool {
	const tooSmallToRepeat = 10
	const halveAmount = 2
	const evenNumber = 2
	if num < tooSmallToRepeat {
		return false
	}

	amountOfDigits := getDigits(num)

	if amountOfDigits%evenNumber != 0 {
		return false
	}

	halvedAmountOfDigits := amountOfDigits / halveAmount
	divisor := int64(1)
	for i := 0; i < halvedAmountOfDigits; i++ {
		divisor *= 10 // We want to create a number like 10, 100, 1000, etc
	}

	firstHalf := num / divisor
	secondHalf := num % divisor

	return firstHalf == secondHalf
}

func getDigits(num int64) int {
	var amountOfDigits int
	temp := num
	for temp > 0 {
		temp /= 10
		amountOfDigits++
	}
	return amountOfDigits
}

func getInvalidIDsP1(ranges []IDRange) []int64 {
	invalidIDs := make([]int64, 0, 100)
	for _, r := range ranges {
		for id := r.start; id <= r.end; id++ {
			if areNumbersRepeating(id) {
				invalidIDs = append(invalidIDs, id)
			}
		}
	}
	return invalidIDs
}

func getInvalidIDsP2(ranges []IDRange) []int64 {
	invalidIDs := make([]int64, 0, 100)
	for _, r := range ranges {
		for id := r.start; id <= r.end; id++ {
			if areNumbersRepeating(id) {
				invalidIDs = append(invalidIDs, id)
			} else if isSameDigit(id) {
				invalidIDs = append(invalidIDs, id)
			} else if isRepeating := checkRepeatingTwoDigitNum(id); isRepeating == true {
				invalidIDs = append(invalidIDs, id)
			} else if isRepeating := checkRepeatingThreeDigitNum(id, 0, getDigits(id)); isRepeating == true {
				invalidIDs = append(invalidIDs, id)
			}
		}
	}
	return invalidIDs
}

func checkRepeatingTwoDigitNum(num int64) bool {
	if num > 100 {
		amountOfDigits := getDigits(num)
		temp := num / 100
		tempDigits := amountOfDigits / 2
		lastDigits := num % 100
		if lastDigits <= 1 {
			return false
		}
		for i := 0; i < tempDigits; i++ {
			if temp%lastDigits != 0 {
				break
			}
			if temp == lastDigits {
				return true
			}
			temp /= 100
			if temp == 0 {
				return false
			}
		}
	}
	return false
}

func checkRepeatingThreeDigitNum(num int64, temp int64, amountOfDigits int) bool {
	temp = num / 1000
	tempDigits := amountOfDigits - 3
	lastDigits := num % 1000
	if lastDigits <= 1 {
		return false
	}
	for i := 0; i < tempDigits; i++ {
		if temp%lastDigits != 0 {
			break
		}
		tempDigits -= 3
		temp /= 1000
		if tempDigits == 0 && temp == 0 {
			return true
		}
	}
	return false
}

func isSameDigit(num int64) bool {
	if num < 0 {
		num = -num // Handle negative numbers
	}
	if num < 10 {
		return true // Single digits always consist of the same digit
	}

	lastDigit := num % 10
	for num > 0 {
		if num%10 != lastDigit {
			return false
		}
		num /= 10
	}
	return true
}

func SumInvalidIDs(input string, isPartOne bool) int64 {
	ranges, err := parseIDRanges(input)
	if err != nil {
		return 0
	}
	var invalidIDs []int64
	if isPartOne {
		invalidIDs = getInvalidIDsP1(ranges)
	} else {
		invalidIDs = getInvalidIDsP2(ranges)
	}
	var sum int64
	for _, id := range invalidIDs {
		sum += id
	}
	return sum
}
