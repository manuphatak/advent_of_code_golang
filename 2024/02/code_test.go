package main

import (
	"testing"
)

func TestIsSafePart2(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "7 6 4 2 1", want: true},
		{input: "1 2 7 8 9", want: false},
		{input: "9 7 6 2 1", want: false},
		{input: "1 3 2 4 5", want: true},
		{input: "8 6 4 4 1", want: true},
		{input: "1 3 6 7 9", want: true},
		{input: "1 5 6 7 9", want: true},
		{input: "4 3 6 7 9", want: true},
		{input: "48 51 54 56 60", want: true},
		{input: "1 2 3 0 4", want: true},
		{input: "32 32 35 32 30", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			report := parseReport(tt.input)
			if got := isSafePart2(report); got != tt.want {
				t.Errorf("isSafePart2(%v) = %v, want %v", report, got, tt.want)
			}
		})
	}
}
