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
