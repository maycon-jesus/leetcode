package main

import (
	"testing"
)

func TestMostBooked(t *testing.T) {
	tests := []struct {
		n        int
		meetings [][]int
		want     int
	}{
		{
			n:        2,
			meetings: [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}},
			want:     0,
		},
		{
			n:        3,
			meetings: [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}},
			want:     1,
		},
	}

	for i, tt := range tests {
		got := mostBooked(tt.n, tt.meetings)
		if got != tt.want {
			t.Errorf("Test %d: mostBooked(%d, %v) = %d, want %d", i+1, tt.n, tt.meetings, got, tt.want)
		}
	}
}
