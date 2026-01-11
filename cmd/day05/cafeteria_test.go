package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantRanges [][]int64
		wantIDs    []int64
	}{
		{
			name:       "multiple lines",
			input:      "10-20\n30-40\n\n15\n25\n35\n45",
			wantRanges: [][]int64{{10, 20}, {30, 40}},
			wantIDs:    []int64{15, 25, 35, 45},
		},
		{
			name:       "empty input",
			input:      "",
			wantRanges: [][]int64(nil),
			wantIDs:    []int64(nil),
		},
		{
			name:       "only ranges",
			input:      "1-5\n10-15",
			wantRanges: [][]int64{{1, 5}, {10, 15}},
			wantIDs:    []int64(nil),
		},
		{
			name:       "only ids",
			input:      "\n\n100\n200",
			wantRanges: [][]int64(nil),
			wantIDs:    []int64{100, 200},
		},
		{
			name:       "single range and single id",
			input:      "5-10\n\n42",
			wantRanges: [][]int64{{5, 10}},
			wantIDs:    []int64{42},
		},
		{
			name:       "extra whitespace",
			input:      "  10-20  \n  30-40  \n\n  15  \n  25  ",
			wantRanges: [][]int64{{10, 20}, {30, 40}},
			wantIDs:    []int64{15, 25},
		},
		{
			name:       "invalid range format ignored",
			input:      "10-20\ninvalid\n30-40\n\n15",
			wantRanges: [][]int64{{10, 20}, {30, 40}},
			wantIDs:    []int64{15},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotRanges, gotIDs := parseInput(test.input)
			if !reflect.DeepEqual(gotRanges, test.wantRanges) {
				t.Errorf("parseInput(%q)\ngot  = %v\nwant = %v", test.input, gotRanges, test.wantRanges)
			}
			if !reflect.DeepEqual(gotIDs, test.wantIDs) {
				t.Errorf("parseInput(%q)\ngot  = %v\nwant = %v", test.input, gotIDs, test.wantIDs)
			}
		})
	}
}

func TestCountValidIDs(t *testing.T) {
	tests := []struct {
		name        string
		inputRanges [][]int64
		inputIDs    []int64
		want        int
	}{
		{
			name:        "multiple lines",
			inputRanges: [][]int64{{10, 20}, {30, 40}},
			inputIDs:    []int64{15, 25, 35, 45},
			want:        2,
		},
		{
			name:        "no ids",
			inputRanges: [][]int64{{10, 20}},
			inputIDs:    []int64{},
			want:        0,
		},
		{
			name:        "no ranges",
			inputRanges: [][]int64{},
			inputIDs:    []int64{10, 20},
			want:        0,
		},
		{
			name:        "ids at boundaries",
			inputRanges: [][]int64{{10, 20}},
			inputIDs:    []int64{10, 20, 9, 21},
			want:        2,
		},
		{
			name:        "overlapping ranges",
			inputRanges: [][]int64{{10, 20}, {15, 25}},
			inputIDs:    []int64{12, 18, 22},
			want:        3,
		},
		{
			name:        "all ids valid",
			inputRanges: [][]int64{{0, 100}},
			inputIDs:    []int64{0, 50, 100},
			want:        3,
		},
		{
			name:        "no ids valid",
			inputRanges: [][]int64{{10, 20}, {30, 40}},
			inputIDs:    []int64{5, 25, 45},
			want:        0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := countValidIDs(test.inputIDs, test.inputRanges)
			if got != test.want {
				t.Errorf("countValidIDs(%v, %v) = %d; want %d", test.inputIDs, test.inputRanges, got, test.want)
			}
		})
	}
}

func TestCountValidIDsInRange(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int64
		want  int64
	}{
		{
			name:  "multiple lines",
			input: [][]int64{{10, 20}, {30, 40}},
			want:  22,
		},
		{
			name:  "single range",
			input: [][]int64{{5, 10}},
			want:  6,
		},
		{
			name:  "overlapping ranges",
			input: [][]int64{{10, 20}, {15, 25}},
			want:  16,
		},
		{
			name:  "fully contained range",
			input: [][]int64{{10, 50}, {20, 30}},
			want:  41,
		},
		{
			name:  "disjoint ranges",
			input: [][]int64{{1, 5}, {10, 15}},
			want:  11,
		},
		{
			name:  "adjacent ranges",
			input: [][]int64{{1, 5}, {6, 10}},
			want:  10,
		},
		{
			name:  "unsorted input ranges",
			input: [][]int64{{30, 40}, {10, 20}},
			want:  22,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := countValidIDsInRange(test.input)
			if got != test.want {
				t.Errorf("countValidIDsInRange(%v) = %d; want %d", test.input, got, test.want)
			}
		})
	}
}
