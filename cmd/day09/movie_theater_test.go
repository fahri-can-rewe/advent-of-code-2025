package main

import "testing"

func TestParseInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][]int
	}{
		{
			name: "Basic Test",
			input: `
			1,2
			3,4
			5,6
			`,
			want: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name: "Single Coordinate",
			input: `
			7,8
			`,
			want: [][]int{{7, 8}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseInput(tt.input)
			if len(result) != len(tt.want) {
				t.Fatalf("want length %d, got %d", len(tt.want), len(result))
			}
		})
	}
}

func TestFindLargestRectangle(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name: "Basic Test",
			input: `
			1,2
			3,4
			5,6
			`,
			want: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coordinates := parseInput(tt.input)
			result := findLargestRectangle(coordinates)
			if result != tt.want {
				t.Errorf("want %d, got %d", tt.want, result)
			}
		})
	}
}
