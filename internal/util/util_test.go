package util

import (
	"testing"
)

func TestReadInput(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    string
		wantErr bool
	}{
		{
			name:    "success",
			path:    "testdata/test_input.txt",
			want:    "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82",
			wantErr: false,
		},
		{
			name:    "file not found",
			path:    "testdata/non-existent-file.txt",
			want:    "",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ReadInput(test.path)
			if (err != nil) != test.wantErr {
				t.Errorf("ReadInput() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("ReadInput() = %q, want %q", got, test.want)
			}
		})
	}
}
