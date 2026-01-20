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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInput(tt.input)
			if len(got) != len(tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
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
			got := useDFS(racks)
			if got != tt.want {
				t.Errorf("useDFS() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
