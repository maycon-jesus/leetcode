package main

import "testing"

func TestPossibleStringCount(t *testing.T) {
	tests := []struct {
		word     string
		expected int
	}{
		{"abbcccc", 5},
		{"abcd", 1},
		{"aaaa", 4},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			result := possibleStringCount(tt.word)
			if result != tt.expected {
				t.Errorf("possibleStringCount(%q) = %d; want %d", tt.word, result, tt.expected)
			}
		})
	}
}
