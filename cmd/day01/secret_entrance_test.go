package main

import (
	"reflect"
	"testing"
)

func TestGenerateInstructions(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Rotator
	}{
		{
			name:  "single instruction R",
			input: "R12",
			want:  []Rotator{{"R", 12}},
		},
		{
			name:  "single instruction L",
			input: "L99",
			want:  []Rotator{{"L", 99}},
		},
		{
			name:  "multiple instructions",
			input: "R12\nL99\nR48",
			want: []Rotator{
				{"R", 12},
				{"L", 99},
				{"R", 48},
			},
		},
		{
			name:  "empty input",
			input: "",
			want:  []Rotator{},
		},
		{
			name:  "input with extra whitespace and empty lines",
			input: "R12\n\nL99\n  ",
			want: []Rotator{
				{"R", 12},
				{"L", 99},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := generateInstructions(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("generateInstructions(%q)\ngot  = %v\nwant = %v", test.input, got, test.want)
			}
		})
	}
}

func TestMoveDialP1(t *testing.T) {
	tests := []struct {
		name  string
		input []Rotator
		want  int
	}{
		{
			name: "no zero hits",
			input: []Rotator{
				{"R", 10},
				{"L", 5},
			},
			want: 0,
		},
		{
			name: "hit zero exactly moving right",
			input: []Rotator{
				{"R", 50},
			},
			want: 1,
		},
		{
			name: "hit zero moving left",
			input: []Rotator{
				{"L", 50},
			},
			want: 1,
		},
		{
			name: "multiple hits and moves",
			input: []Rotator{
				{"R", 50},
				{"L", 50},
				{"L", 50},
			},
			want: 2,
		},
		{
			name: "large steps wrapping around multiple times",
			input: []Rotator{
				{"R", 150},
			},
			want: 1,
		},
		{
			name: "complex sequence wrapping around",
			input: []Rotator{
				{"R", 70},  // 50 + 70 = 120 -> 120-1-99 = 20
				{"L", 40},  // 20 - 40 = -20 -> -20+1+99 = 80
				{"R", 20},  // 80 + 20 = 100 -> 100-1-99 = 0 (hit 1)
				{"R", 100}, // 0 + 100%100 (0) = 0 (hit 2)
			},
			want: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := moveDialP1(test.input)
			if got != test.want {
				t.Errorf("%s: moveDialP1() = %d; want %d", test.name, got, test.want)
			}
		})
	}
}
