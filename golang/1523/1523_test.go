package main

import "testing"

func TestCountOdds(t *testing.T) {
	tests := []struct {
		name string
		low  int
		high int
		want int
	}{
		{
			name: "Example 1: both odd boundaries",
			low:  3,
			high: 7,
			want: 3,
		},
		{
			name: "Example 2: both even boundaries",
			low:  8,
			high: 10,
			want: 1,
		},
		{
			name: "Single number - odd",
			low:  5,
			high: 5,
			want: 1,
		},
		{
			name: "Single number - even",
			low:  4,
			high: 4,
			want: 0,
		},
		{
			name: "Low even, high odd",
			low:  2,
			high: 7,
			want: 3,
		},
		{
			name: "Low odd, high even",
			low:  1,
			high: 6,
			want: 3,
		},
		{
			name: "Consecutive numbers - even to odd",
			low:  4,
			high: 5,
			want: 1,
		},
		{
			name: "Consecutive numbers - odd to even",
			low:  5,
			high: 6,
			want: 1,
		},
		{
			name: "Large range with both even",
			low:  0,
			high: 100,
			want: 50,
		},
		{
			name: "Large range with both odd",
			low:  1,
			high: 99,
			want: 50,
		},
		{
			name: "Zero to small even",
			low:  0,
			high: 10,
			want: 5,
		},
		{
			name: "Zero to small odd",
			low:  0,
			high: 9,
			want: 5,
		},
		{
			name: "Very large range",
			low:  1000000,
			high: 2000000,
			want: 500000,
		},
		{
			name: "Edge case: maximum constraint",
			low:  0,
			high: 1000000000,
			want: 500000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countOdds(tt.low, tt.high)
			if got != tt.want {
				t.Errorf("countOdds(%d, %d) = %d, want %d", tt.low, tt.high, got, tt.want)
			}
		})
	}
}
