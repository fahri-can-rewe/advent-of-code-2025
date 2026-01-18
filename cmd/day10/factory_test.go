package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Machine
	}{
		{
			name:  "single line",
			input: `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}`,
			want: []Machine{
				{
					Target: []bool{false, true, true, false},
					Buttons: [][]int{
						{3},
						{1, 3},
						{2},
						{2, 3},
						{0, 2},
						{0, 1},
					},
				},
			},
		},
		{
			name:  "multiple machines",
			input: "[.#] (0) (1)\n[#.] (1)",
			want: []Machine{
				{
					Target:  []bool{false, true},
					Buttons: [][]int{{0}, {1}},
				},
				{
					Target:  []bool{true, false},
					Buttons: [][]int{{1}},
				},
			},
		},
		{
			name:  "empty input",
			input: "",
			want:  nil,
		},
		{
			name:  "machine with no buttons",
			input: "[#] {10}",
			want: []Machine{
				{
					Target:  []bool{true},
					Buttons: [][]int{},
				},
			},
		},
		{
			name:  "extra spaces and newlines",
			input: "\n\n  [.##.] (0,1)  \n\n  ",
			want: []Machine{
				{
					Target:  []bool{false, true, true, false},
					Buttons: [][]int{{0, 1}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseInput(tt.input)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("parseInput() = %+v, want %+v", result, tt.want)
			}
		})
	}
}

func TestFindMinPresses(t *testing.T) {
	tests := []struct {
		name  string
		input Machine
		want  int
	}{
		{
			name: "single line",
			input: Machine{
				Target: []bool{false, true, true, false},
				Buttons: [][]int{
					{3},
					{1, 3},
					{2},
					{2, 3},
					{0, 2},
					{0, 1},
				},
			},
			want: 2,
		},
		{
			name: "unsolvable configuration",
			input: Machine{
				Target: []bool{true},
				Buttons: [][]int{
					{1}, // toggles out of range light
				},
			},
			want: -1,
		},
		{
			name: "already at target (all off)",
			input: Machine{
				Target:  []bool{false, false},
				Buttons: [][]int{{0}, {1}},
			},
			want: 0,
		},
		{
			name: "multiple ways to solve, choose minimum",
			input: Machine{
				Target: []bool{true, true},
				Buttons: [][]int{
					{0},
					{1},
					{0, 1},
				},
			},
			want: 1, // pressing {0, 1} is 1 press vs pressing {0} and {1} is 2 presses
		},
		{
			name: "no buttons available",
			input: Machine{
				Target:  []bool{true},
				Buttons: [][]int{},
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinPresses(tt.input)
			if result != tt.want {
				t.Errorf("findMinPresses() = %d, want %d", result, tt.want)
			}
		})
	}
}

func TestCountBtnPress(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "single line",
			input: "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
			want:  2,
		},
		{
			name: "full sample input",
			input: "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}\n" +
				"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}\n" +
				"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
			want: 7,
		},
		{
			name:  "mixed solvable and unsolvable machines",
			input: "[.#] (1)\n[#] (1)", // 1st solvable (light 1 is on, button 0 toggles light 1), 2nd unsolvable (light 0 is on, button 0 toggles light 1)
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countBtnPress(tt.input)
			if result != tt.want {
				t.Errorf("countBtnPress() = %d, want %d", result, tt.want)
			}
		})
	}
}
