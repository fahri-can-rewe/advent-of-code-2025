package main

import (
	"reflect"
	"testing"
)

func TestParseIDRanges(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []IDRange
		wantErr bool
	}{
		{
			name:  "small ranges",
			input: "1-3,5-7",
			want: []IDRange{
				{1, 3},
				{5, 7},
			},
			wantErr: false,
		},
		{
			name:    "invalid range format",
			input:   "1-3,5-7,a-b",
			want:    nil,
			wantErr: true,
		},
		{
			name:  "single range",
			input: "10-20",
			want: []IDRange{
				{10, 20},
			},
			wantErr: false,
		},
		{
			name:    "missing end",
			input:   "10-",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "missing start",
			input:   "-20",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "start greater than end",
			input:   "50-40",
			want:    nil,
			wantErr: true,
		},
		{
			name:  "whitespace handling",
			input: " 1-3 ,  5-7  ",
			want: []IDRange{
				{1, 3},
				{5, 7},
			},
			wantErr: false,
		},
		{
			name:    "invalid range format: random text",
			input:   "abc-def",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid range format: missing hyphen",
			input:   "12345",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid range format: double hyphen",
			input:   "10--20",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid range format: extra characters",
			input:   "10-20abc",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid range format: empty part",
			input:   "10-20,,30-40",
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseIDRanges(test.input)
			if (err != nil) != test.wantErr {
				t.Errorf("%s: parseIDRanges() error = %v, wantErr %v", test.name, err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%s: parseIDRanges() got = %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestIsInvalidID(t *testing.T) {
	tests := []struct {
		name string
		num  int64
		want bool
	}{
		{
			name: "non-repeating two digit number",
			num:  12,
			want: false,
		},
		{
			name: "repeating two digit number",
			num:  11,
			want: true,
		},
		{
			name: "repeating big two digit number",
			num:  99,
			want: true,
		},
		{
			name: "repeating three digit number",
			num:  111,
			want: true,
		},
		{
			name: "repeating four digit number",
			num:  1010,
			want: true,
		},
		{
			name: "non-repeating four digit number",
			num:  9990,
			want: false,
		},
		{
			name: "repeating six digit number",
			num:  446446,
			want: true,
		},
		{
			name: "zero",
			num:  0,
			want: false,
		},
		{
			name: "single digit",
			num:  7,
			want: false,
		},
		{
			name: "negative repeating number",
			num:  -11,
			want: false,
		},
		{
			name: "three digits almost repeating",
			num:  121,
			want: false,
		},
		{
			name: "large repeating number",
			num:  123456789123456789,
			want: true,
		},
		{
			name: "large non-repeating number",
			num:  123456789987654321,
			want: false,
		},
		{
			name: "halves differ by one digit",
			num:  1213,
			want: false,
		},
		{
			name: "halves are reversed",
			num:  1221,
			want: false,
		},
		{
			name: "repeating but odd length",
			num:  12121,
			want: false,
		},
		{
			name: "all same digits four",
			num:  1111,
			want: true,
		},
		{
			name: "all same digits six",
			num:  111111,
			want: true,
		},
		{
			name: "four digits mirrored but not repeating",
			num:  1001,
			want: false,
		},
		{
			name: "four digits first half same, second half same",
			num:  1100,
			want: false,
		},
		{
			name: "example repeating sequence twice",
			num:  12341234,
			want: true,
		},
		{
			name: "example repeating sequence thrice",
			num:  123123123,
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isInvalidID(test.num)
			if got != test.want {
				t.Errorf("%s: isInvalidID(%d) = %v, want %v", test.name, test.num, got, test.want)
			}
		})
	}
}

func TestGetInvalidIDsP1(t *testing.T) {
	tests := []struct {
		name  string
		input []IDRange
		want  []int64
	}{
		{
			name: "single range",
			input: []IDRange{
				{10, 22},
			},
			want: []int64{11, 22},
		},
		{
			name: "multiple ranges",
			input: []IDRange{
				{10, 25},
				{30, 45},
			},
			want: []int64{11, 22, 33, 44},
		},
		{
			name:  "empty input slice",
			input: []IDRange{},
			want:  []int64{},
		},
		{
			name: "no invalid ids in range",
			input: []IDRange{
				{12, 15},
			},
			want: []int64{},
		},
		{
			name: "single point range with invalid id",
			input: []IDRange{
				{11, 11},
			},
			want: []int64{11},
		},
		{
			name: "single point range with valid id",
			input: []IDRange{
				{12, 12},
			},
			want: []int64{},
		},
		{
			name: "overlapping ranges",
			input: []IDRange{
				{10, 15},
				{11, 22},
			},
			want: []int64{11, 11, 22},
		},
		{
			name: "zero and negative ranges",
			input: []IDRange{
				{-11, 11},
			},
			want: []int64{11},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getInvalidIDsP1(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%s: getInvalidIDsP1(%v) = %v, want %v", test.name, test.input, got, test.want)
			}
		})
	}
}

func TestSumInvalidIDs_P1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "single range",
			input: "10-25",
			want:  33,
		},
		{
			name:  "multiple ranges",
			input: "10-25,30-45",
			want:  110,
		},
		{
			name:  "empty input",
			input: "",
			want:  0,
		},
		{
			name:  "invalid range format",
			input: "abc-def",
			want:  0,
		},
		{
			name:  "start greater than end",
			input: "50-40",
			want:  0,
		},
		{
			name:  "whitespace handling",
			input: " 10-25 ,  30-45 ",
			want:  110,
		},
		{
			name:  "overlapping ranges",
			input: "10-15,11-22",
			want:  44,
		},
		{
			name:  "no invalid IDs",
			input: "12-15",
			want:  0,
		},
		{
			name:  "single point range invalid",
			input: "11-11",
			want:  11,
		},
		{
			name:  "single point range valid",
			input: "12-12",
			want:  0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := SumInvalidIDs(test.input, true)
			if got != test.want {
				t.Errorf("%s: SumInvalidIDs(%q) = %v, want %v", test.name, test.input, got, test.want)
			}
		})
	}
}

func TestSumInvalidIDs_P2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "single range",
			input: "11-22",
			want:  33,
		},
		{
			name:  "multiple ranges",
			input: "95-115,998-1012",
			want:  2219,
		},
		{
			name:  "multiple ranges",
			input: "1188511880-1188511890",
			want:  1188511885,
		},
		{
			name:  "multiple ranges",
			input: "565653-565659",
			want:  565656,
		},
		{
			name:  "multiple ranges",
			input: "824824821-824824827",
			want:  824824824,
		},
		{
			name:  "multiple ranges",
			input: "2121212118-2121212124",
			want:  2121212121,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := SumInvalidIDs(test.input, false)
			if got != test.want {
				t.Errorf("%s: SumInvalidIDs(%q) = %v, want %v", test.name, test.input, got, test.want)
			}
		})
	}
}
