package main

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Example 1", "42", 42},
		{"Example 2", "   -042", -42},
		{"Example 3", "1337c0d3", 1337},
		{"Example 4", "0-1", 0},
		{"Example 5", "words and 987", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myAtoi(tt.input)
			if got != tt.want {
				t.Errorf("myAtoi(%q) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}
