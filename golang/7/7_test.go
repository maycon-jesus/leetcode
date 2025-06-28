package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Example 1", 123, 321},
		{"Example 2", -123, -321},
		{"Example 3", 120, 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.input)
			if got != tt.want {
				t.Errorf("reverse(%d) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}
