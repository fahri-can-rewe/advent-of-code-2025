package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		wantNums       [][]int
		wantOperations []string
	}{
		{
			name:           "duo lines",
			input:          "123 328  51 64\n 45 64  387 23\n*   +   *   +",
			wantNums:       [][]int{{123, 328, 51, 64}, {45, 64, 387, 23}},
			wantOperations: []string{"*", "+", "*", "+"},
		},
		{
			name:           "empty input",
			input:          "",
			wantNums:       [][]int(nil),
			wantOperations: []string(nil),
		},
		{
			name:           "single line numbers",
			input:          "1 2 3",
			wantNums:       [][]int{{1, 2, 3}},
			wantOperations: []string(nil),
		},
		{
			name:           "single line operations",
			input:          "+ * +",
			wantNums:       [][]int(nil),
			wantOperations: []string{"+", "*", "+"},
		},
		{
			name:           "mixed non-numeric characters",
			input:          "10, 20; 30\n# comment\n+ - * /",
			wantNums:       [][]int{{10, 20, 30}},
			wantOperations: []string{"+", "*"},
		},
		{
			name:           "numbers and operations on same line",
			input:          "1 + 2 * 3",
			wantNums:       [][]int{{1, 2, 3}},
			wantOperations: []string{"+", "*"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotNums, gotOperations := parseInput(test.input)
			if !reflect.DeepEqual(gotNums, test.wantNums) {
				t.Errorf("parseInput(%q) got nums = %v; want %v", test.input, gotNums, test.wantNums)
			}
			if !reflect.DeepEqual(gotOperations, test.wantOperations) {
				t.Errorf("parseInput(%q) got operations = %v; want %v", test.input, gotOperations, test.wantOperations)
			}
		})
	}
}

func TestSumUpCalculations(t *testing.T) {
	tests := []struct {
		name      string
		inputNums [][]int
		inputOps  []string
		want      int64
	}{
		{
			name:      "triple line input",
			inputNums: [][]int{{123, 328, 51, 64}, {45, 64, 387, 23}, {6, 98, 215, 314}},
			inputOps:  []string{"*", "+", "*", "+"},
			want:      4277556,
		},
		{
			name:      "empty input",
			inputNums: [][]int(nil),
			inputOps:  []string(nil),
			want:      0,
		},
		{
			name:      "single addition",
			inputNums: [][]int{{10}, {20}},
			inputOps:  []string{"+"},
			want:      30,
		},
		{
			name:      "single multiplication",
			inputNums: [][]int{{10}, {20}},
			inputOps:  []string{"*"},
			want:      200,
		},
		{
			name:      "no operations",
			inputNums: [][]int{{1, 2}, {3, 4}},
			inputOps:  []string(nil),
			want:      0,
		},
		{
			name:      "mixed operations",
			inputNums: [][]int{{2, 3}, {4, 5}},
			inputOps:  []string{"+", "*"},
			want:      21,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sumUpCalculations(test.inputNums, test.inputOps)
			if got != test.want {
				t.Errorf("sumUpCalculations(%v, %v) = %v; want %v", test.inputNums, test.inputOps, got, test.want)
			}
		})
	}
}
