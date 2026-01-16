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
		{
			name: "Negative Coordinates",
			input: `
			-1,-2
			0,0
			`,
			want: [][]int{{-1, -2}, {0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseInput(tt.input)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("parseInput() = %v, want %v", result, tt.want)
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
		{
			name: "Same Points",
			input: `
			1,1
			1,1
			`,
			want: 1,
		},
		{
			name: "Horizontal Line",
			input: `
			1,1
			5,1
			`,
			want: 5,
		},
		{
			name: "Vertical Line",
			input: `
			1,1
			1,5
			`,
			want: 5,
		},
		{
			name: "Negative Range",
			input: `
			-2,-2
			2,2
			`,
			want: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coordinates := parseInput(tt.input)
			result := findLargestRectangle(coordinates)
			if result != tt.want {
				t.Errorf("findLargestRectangle() = %d, want %d", result, tt.want)
			}
		})
	}
}
