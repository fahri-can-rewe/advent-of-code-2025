package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][]int
	}{
		{
			name:  "single line",
			input: "123456789",
			want:  [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			name:  "duo lines",
			input: "9876543\n23456789",
			want:  [][]int{{9, 8, 7, 6, 5, 4, 3}, {2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			name:  "empty input",
			input: "",
			want:  [][]int{nil},
		},
		{
			name:  "whitespace only",
			input: "  \n  ",
			want:  [][]int{nil},
		},
		{
			name:  "single digit",
			input: "5",
			want:  [][]int{{5}},
		},
		{
			name:  "different line lengths",
			input: "12\n1234\n1",
			want:  [][]int{{1, 2}, {1, 2, 3, 4}, {1}},
		},
		{
			name:  "leading and trailing whitespace",
			input: "  123  \n  456  ",
			want:  [][]int{{1, 2, 3, 0, 0}, {0, 0, 4, 5, 6}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseInput(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("parseInput() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestIsHighestNum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "no nines",
			nums: []int{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "one nine at wrong index",
			nums: []int{1, 2, 3, 9, 5, 8, 3},
			want: false,
		},
		{
			name: "one nine at last index",
			nums: []int{1, 2, 3, 9, 5, 8, 9},
			want: false,
		},
		{
			name: "one nine at first index",
			nums: []int{9, 2, 3, 9, 5, 8, 4},
			want: false,
		},
		{
			name: "nine at first and second index",
			nums: []int{9, 9, 3, 9, 5, 8, 4},
			want: true,
		},
		{
			name: "nine at first, second and last index",
			nums: []int{9, 9, 3, 9, 5, 8, 4, 9},
			want: true,
		},
		{
			name: "nine at second and last index",
			nums: []int{7, 9, 3, 9, 5, 8, 4, 9},
			want: true,
		},
		{
			name: "nine at first and last index",
			nums: []int{9, 7, 3, 9, 5, 8, 4, 9},
			want: true,
		},
		{
			name: "empty slice",
			nums: []int{},
			want: false,
		},
		{
			name: "single element nine",
			nums: []int{9},
			want: false,
		},
		{
			name: "single element not nine",
			nums: []int{1},
			want: false,
		},
		{
			name: "two elements both nine",
			nums: []int{9, 9},
			want: true,
		},
		{
			name: "two elements one nine",
			nums: []int{9, 1},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isHighestNum(test.nums)
			if got != test.want {
				t.Errorf("isHighestNum() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestFindSecondDigit(t *testing.T) {
	tests := []struct {
		name string
		idx  int
		nums []int
		want int
	}{
		{
			name: "find the 9",
			idx:  1,
			nums: []int{9, 3, 5, 6, 9},
			want: 9,
		},
		{
			name: "find the 7",
			idx:  1,
			nums: []int{9, 3, 5, 6, 7, 2, 4},
			want: 7,
		},
		{
			name: "find the 5",
			idx:  4,
			nums: []int{9, 3, 5, 6, 9, 5, 4, 1, 1, 1, 2, 4},
			want: 5,
		},
		{
			name: "nothing after idx",
			idx:  4,
			nums: []int{1, 2, 3, 4, 9},
			want: 1,
		},
		{
			name: "only one digit after idx (larger than 1)",
			idx:  1,
			nums: []int{9, 9, 5},
			want: 5,
		},
		{
			name: "multiple same max digits",
			idx:  0,
			nums: []int{1, 5, 2, 5, 3},
			want: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findSecondDigit(test.idx, test.nums)
			if got != test.want {
				t.Errorf("findSecondDigit(%d, %v) = %v, want %v", test.idx, test.nums, got, test.want)
			}
		})
	}
}

func TestFindTwoLargestDigits(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty slice",
			nums: []int{},
			want: 0,
		},
		{
			name: "single element",
			nums: []int{5},
			want: 5,
		},
		{
			name: "isHighestNum - 99 at start",
			nums: []int{9, 9, 1, 2},
			want: 99,
		},
		{
			name: "isHighestNum - 99 at start and end",
			nums: []int{9, 1, 2, 9},
			want: 99,
		},
		{
			name: "general case - 87",
			nums: []int{8, 3, 5, 7, 2},
			want: 87,
		},
		{
			name: "general case - 98",
			nums: []int{9, 3, 8, 2},
			want: 98,
		},
		{
			name: "largest is before end, second largest is after it - 92",
			nums: []int{1, 8, 9, 2},
			want: 92,
		},
		{
			name: "second largest is before largest - 89",
			nums: []int{1, 8, 3, 9},
			want: 89,
		},
		{
			name: "multiple nines - 99 via logic",
			nums: []int{1, 9, 2, 9, 3},
			want: 99,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findTwoLargestDigits(test.nums)
			if got != test.want {
				t.Errorf("findTwoLargestDigits(%v) = %v, want %v", test.nums, got, test.want)
			}
		})
	}
}

func TestSumUpJoltages(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "whitespace only",
			input: "  \n  ",
			want:  0,
		},
		{
			name:  "single line - 123 -> 23",
			input: "123",
			want:  23,
		},
		{
			name:  "multiple lines",
			input: "123\n456",
			want:  79,
		},
		{
			name:  "multiple lines with varying lengths",
			input: "9876543\n23456789\n5",
			want:  192,
		},
		{
			name:  "leading and trailing whitespace",
			input: "  123  \n  456  ",
			want:  87,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sumUpJoltages(test.input)
			if got != test.want {
				t.Errorf("sumUpJoltages(%q) = %v, want %v", test.input, got, test.want)
			}
		})
	}
}
