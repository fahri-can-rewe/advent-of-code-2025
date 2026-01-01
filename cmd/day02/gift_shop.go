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
	const divisibleByTen = 10
	const halveAmount = 2
	const evenNumber = 2
	if num < tooSmallToRepeat {
		return false
	}

	var amountOfDigits int
	temp := num
	for temp > 0 {
		temp /= divisibleByTen
		amountOfDigits++
	}

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

func getInvalidIDs(ranges []IDRange) []int64 {
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

func SumInvalidIDs(input string) int64 {
	ranges, err := parseIDRanges(input)
	if err != nil {
		return 0
	}
	invalidIDs := getInvalidIDs(ranges)
	var sum int64
	for _, id := range invalidIDs {
		sum += id
	}
	return sum
}
