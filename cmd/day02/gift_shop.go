package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const everyRegexMatch = -1

type IDRange struct {
	start int64
	end   int64
}

func parseIDRanges(input string) ([]IDRange, error) {
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
		ranges = append(ranges, IDRange{start: start, end: end})
	}
	return ranges, nil
}

func areNumbersRepeating(num int64) bool {
	s := strconv.FormatInt(num, 10)
	numLen := len(s)

	if numLen%2 != 0 {
		return false
	}

	half := numLen / 2
	return s[:half] == s[half:]
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
