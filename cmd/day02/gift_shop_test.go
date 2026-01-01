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

func TestAreNumbersRepeating(t *testing.T) {
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
			name: "non-repeating three digit number",
			num:  111,
			want: false,
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := areNumbersRepeating(test.num)
			if got != test.want {
				t.Errorf("%s: areNumbersRepeating(%d) = %v, want %v", test.name, test.num, got, test.want)
			}
		})
	}
}
