package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const allPartsOfTheString = 3
const leftNumBeforeDash = 1
const rightNumAfterDash = 2

func parseInput(input string) ([][]int64, []int64) {
	sections := strings.Split(input, "\n\n")
	var idRanges [][]int64
	var validIDs []int64

	rangeRegex := regexp.MustCompile(`(\d+)-(\d+)`)
	idRegex := regexp.MustCompile(`^\d+$`)

	ranges := parseIDRanges(sections, rangeRegex, idRanges)
	ids := parseIDs(sections, idRegex, validIDs)
	return ranges, ids
}

func parseIDs(sections []string, idRegex *regexp.Regexp, validIDs []int64) []int64 {
	if len(sections) > 1 {
		lines := strings.Split(strings.TrimSpace(sections[1]), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if idRegex.MatchString(line) {
				id, _ := strconv.ParseInt(line, 10, 64)
				validIDs = append(validIDs, id)
			}
		}
	}
	return validIDs
}

func parseIDRanges(sections []string, rangeRegex *regexp.Regexp, idRanges [][]int64) [][]int64 {
	if len(sections) > 0 {
		lines := strings.Split(strings.TrimSpace(sections[0]), "\n")
		for _, line := range lines {
			matches := rangeRegex.FindStringSubmatch(line)
			if len(matches) == allPartsOfTheString {
				start, _ := strconv.ParseInt(matches[leftNumBeforeDash], 10, 64)
				end, _ := strconv.ParseInt(matches[rightNumAfterDash], 10, 64)
				idRanges = append(idRanges, []int64{start, end})
			}
		}
	}
	return idRanges
}

func countValidIDs(validIDs []int64, idRanges [][]int64) int {
	counter := 0
	for _, id := range validIDs {
		for _, r := range idRanges {
			if id >= r[0] && id <= r[1] {
				counter++
				break
			}
		}
	}
	return counter
}

func countValidIDsInRange(idRanges [][]int64) int64 {
	// 1. Sort ranges by start value
	sort.Slice(idRanges, func(i, j int) bool {
		return idRanges[i][0] < idRanges[j][0]
	})

	// 2. Merge overlapping ranges
	merged := [][]int64{idRanges[0]}
	for i := 1; i < len(idRanges); i++ {
		last := &merged[len(merged)-1]
		current := idRanges[i]

		if current[0] <= (*last)[1] {
			// If current range overlaps or touches the last merged range, merge them
			if current[1] > (*last)[1] {
				(*last)[1] = current[1]
			}
		} else {
			// Otherwise, add it as a new disjoint range
			merged = append(merged, current)
		}
	}

	// 3. Sum the lengths of the merged ranges
	var totalFreshIDs int64
	for _, r := range merged {
		totalFreshIDs += r[1] - r[0] + 1
	}

	return totalFreshIDs
}
