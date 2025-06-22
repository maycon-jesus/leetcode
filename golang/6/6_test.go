package main

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		s        string
		numRows  int
		expected string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
	}

	for i, tt := range tests {
		result := convert(tt.s, tt.numRows)
		if result != tt.expected {
			t.Errorf("Test %d failed: convert(%q, %d) = %q; want %q", i+1, tt.s, tt.numRows, result, tt.expected)
		}
	}
}
