package main

import "testing"

func TestParseInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string][]string
	}{
		{
			name:  "single line",
			input: "aaa: you hhh",
			want: map[string][]string{
				"aaa": {"you", "hhh"},
			},
		},
		{
			name: "multiple lines",
			input: `aaa: you hhh
					you: bbb ccc
					bbb: ddd eee
					ccc: ddd eee fff
					ddd: ggg`,
			want: map[string][]string{
				"aaa": {"you", "hhh"},
				"you": {"bbb", "ccc"},
				"bbb": {"ddd", "eee"},
				"ccc": {"ddd", "eee", "fff"},
				"ddd": {"ggg"},
			},
		},
		{
			name:  "extra whitespace and multiple spaces between outputs",
			input: "  aaa:   you    hhh   ",
			want: map[string][]string{
				"aaa": {"you", "hhh"},
			},
		},
		{
			name:  "key with no outputs",
			input: "aaa:",
			want: map[string][]string{
				"aaa": {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInput(tt.input)
			if len(got) != len(tt.want) {
				t.Errorf("parseInput() length = %d, want %d (got: %v)", len(got), len(tt.want), got)
				return
			}
		})
	}
}

func TestUseDFS(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "minimal test (2 paths)",
			input: `you: bbb ccc
					bbb: out
					ccc: out`,
			want: 2,
		},
		{
			name: "sample from description",
			input: `aaa: you hhh
					you: bbb ccc
					bbb: ddd eee
					ccc: ddd eee fff
					ddd: ggg
					eee: out
					fff: out
					ggg: out
					hhh: ccc fff iii
					iii: out`,
			want: 5,
		},
		{
			name: "no paths to out",
			input: `you: bbb
					bbb: ccc
					ccc: ddd`,
			want: 0,
		},
		{
			name:  "direct path to out",
			input: `you: out`,
			want:  1,
		},
		{
			name: "diamond shape",
			input: `you: a b
					a: c
					b: c
					c: out`,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			racks := parseInput(tt.input)
			got := useDFS(racks, false)
			if got != tt.want {
				t.Errorf("Part 1 useDFS() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestCountPathsWithTargets(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "expect 1 count",
			input: `aaa: you hhh
					you: bbb ccc
					svr: ddd ccc
					ccc: fft
					fft: eee
					eee: dac
					dac: ggg
					ggg: out`,
			want: 1,
		},
		{
			name: "expect 0 count",
			input: `aaa: you hhh
					you: bbb ccc
					svr: ddd ccc
					ccc: fft
					fft: eee
					eee: dac
					ggg: out`,
			want: 0,
		},
		{
			name: "reverse order (dac then fft)",
			input: `svr: a
					a: dac
					dac: b
					b: fft
					fft: out`,
			want: 1,
		},
		{
			name: "multiple paths with both targets",
			input: `svr: a b
					a: fft
					b: fft
					fft: dac
					dac: out`,
			want: 2,
		},
		{
			name: "one target missing",
			input: `svr: a
					a: fft
					fft: out`,
			want: 0,
		},
		{
			name: "targets visited but no path to out",
			input: `svr: a
					a: fft
					fft: dac
					dac: deadend`,
			want: 0,
		},
		{
			name: "targets are start and end",
			input: `svr: out
					svr: dac
					dac: fft
					fft: out`,
			want: 1,
		},
		{
			name: "complex diamond with targets",
			input: `svr: a b
					a: fft
					b: fft
					fft: c d
					c: dac
					d: dac
					dac: out`,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			racks := parseInput(tt.input)
			got := useDFS(racks, true)
			if got != tt.want {
				t.Errorf("Part 2 useDFS() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
