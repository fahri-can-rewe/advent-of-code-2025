package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][]rune
	}{
		{
			name:  "single line",
			input: ".......S.......",
			want: [][]rune{
				[]rune(".......S......."),
			},
		},
		{
			name:  "multiple lines",
			input: ".......S.......\n...............\n.......^.......\n...............",
			want: [][]rune{
				[]rune(".......S......."),
				[]rune("..............."),
				[]rune(".......^......."),
				[]rune("..............."),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseInput(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("parseInput(%q) = %v, want %v", test.input, got, test.want)
			}
		})
	}
}

func TestFindS(t *testing.T) {
	tests := []struct {
		name string
		grid [][]rune
		want int
	}{
		{
			name: "S in the middle",
			grid: [][]rune{
				[]rune(".......S......."),
			},
			want: 7,
		},
		{
			name: "S at the start",
			grid: [][]rune{
				[]rune("S......"),
			},
			want: 0,
		},
		{
			name: "S at the end",
			grid: [][]rune{
				[]rune("...........S"),
			},
			want: 11,
		},
		{
			name: "S at the beginning",
			grid: [][]rune{
				[]rune("...S........."),
			},
			want: 3,
		},
		{
			name: "S at the towards end",
			grid: [][]rune{
				[]rune("..........S.."),
			},
			want: 10,
		},
		{
			name: "mutiple lines",
			grid: [][]rune{
				[]rune("...........S."),
				[]rune("............."),
				[]rune("............."),
			},
			want: 11,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findS(test.grid)
			if got != test.want {
				t.Errorf("findS(%v) = %d, want %d", test.grid, got, test.want)
			}
		})
	}
}

func TestCountBeamSplitters(t *testing.T) {
	tests := []struct {
		name string
		grid [][]rune
		want int
	}{
		{
			name: "no splitters",
			grid: [][]rune{
				[]rune(".......S......."),
				[]rune("..............."),
				[]rune("..............."),
			},
			want: 0,
		},
		{
			name: "one splitter",
			grid: [][]rune{
				[]rune(".......S......."),
				[]rune("..............."),
				[]rune(".......^......."),
			},
			want: 1,
		},
		{
			name: "multiple splitters but no hits",
			grid: [][]rune{
				[]rune(".......S......."),
				[]rune("..............."),
				[]rune("......^.^......"),
				[]rune("..............."),
				[]rune(".....^...^....."),
			},
			want: 0,
		},
		{
			name: "splitters at edges",
			grid: [][]rune{
				[]rune("S.............."),
				[]rune("^.............^"),
				[]rune("..............."),
			},
			want: 1,
		},
		{
			name: "multiple splitters multiple hits",
			grid: [][]rune{
				[]rune(".......S......."),
				[]rune(".......^......."),
				[]rune("......^.^......"),
				[]rune("..............."),
				[]rune(".....^...^....."),
			},
			want: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := countBeamSplitters(test.grid)
			if got != test.want {
				t.Errorf("countBeamSplitters(%v) = %d, want %d", test.grid, got, test.want)
			}
		})
	}
}

func TestCountTimelines(t *testing.T) {
	sampleInput := `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`
	diagram := parseInput(sampleInput)
	want := int64(40)
	got := countTimelines(diagram)
	if got != want {
		t.Errorf("countTimelines(sample) = %d, want %d", got, want)
	}
}
