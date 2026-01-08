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

func TestDelAsMuchPaperAsPossible(t *testing.T) {
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
	got := delAsMuchPaperAsPossible(grid)
	want := 43

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
