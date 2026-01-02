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
	s := strconv.FormatInt(num, 10)
	n := len(s)
	if n < 2 || n%2 != 0 {
		return false
	}
	half := n / 2
	return s[:half] == s[half:]
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

func isInvalidID(num int64) bool {
	textNum := strconv.FormatInt(num, 10)
	amountOfDigits := len(textNum)
	if amountOfDigits < 2 {
		return false
	}
	// Try all possible lengths of the repeating sequence
	for l := 1; l <= amountOfDigits/2; l++ {
		if amountOfDigits%l == 0 {
			pattern := textNum[:l]
			isRepeating := true
			for i := l; i < amountOfDigits; i += l {
				if textNum[i:i+l] != pattern {
					isRepeating = false
					break
				}
			}
			if isRepeating {
				return true
			}
		}
	}
	return false
}

func getInvalidIDsP2(ranges []IDRange) []int64 {
	invalidIDs := make([]int64, 0, 100)
	for _, r := range ranges {
		for id := r.start; id <= r.end; id++ {
			if isInvalidID(id) {
				invalidIDs = append(invalidIDs, id)
			}
		}
	}
	return invalidIDs
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
