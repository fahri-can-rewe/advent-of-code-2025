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
