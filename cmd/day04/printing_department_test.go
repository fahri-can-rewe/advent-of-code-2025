package main

import (
	"testing"
)

func TestCountForkliftAccess(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`
	grid := parseInput(input)
	got := countForkliftAccess(grid)
	want := 13

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
